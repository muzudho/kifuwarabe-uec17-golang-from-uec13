package position

import (
	"math/rand"
	"os"

	// Entities

	ren "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/ren"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamerecord"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models/color"
)

// Position - 盤
type Position struct {
	// 盤
	board []color.Color
	// 呼吸点を数えるための一時盤
	checkBoard []int
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ models.Point
	// MovesNum - 手数
	MovesNum int
	// Record - 棋譜
	Record []*gamerecord.GameRecordItem
	// 二重ループ
	iteratorWithoutWall func(func(models.Point))
	// UCT計算中の子の数
	uctChildrenSize int
}

// TemporaryPosition - 盤をコピーするときの一時メモリーとして使います
type TemporaryPosition struct {
	// 盤
	Board []color.Color
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ models.Point
}

// CopyPosition - 盤データのコピー。
func (position1 *Position) CopyPosition(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel) *TemporaryPosition {
	var temp = new(TemporaryPosition)
	temp.Board = make([]color.Color, readonlyGameSettingsModel.GetSentinelBoardArea())
	copy(temp.Board[:], position1.board[:])
	temp.KoZ = position1.KoZ
	return temp
}

// ImportPosition - 盤データのコピー。
func (position1 *Position) ImportPosition(temp *TemporaryPosition) {
	copy(position1.board[:], temp.Board[:])
	position1.KoZ = temp.KoZ
}

// NewPosition - 空っぽの局面を生成します
// あとで InitPosition() を呼び出してください
func NewPosition() *Position {
	return new(Position)
}

// InitPosition - 局面の初期化。
func (position1 *Position) InitPosition(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel) {
	position1.Record = make([]*gamerecord.GameRecordItem, readonlyGameSettingsModel.GetMaxMovesNum())
	position1.uctChildrenSize = readonlyGameSettingsModel.GetBoardArea() + 1

	// サイズが変わっているケースに対応するため、配列の作り直し
	var boardMax = readonlyGameSettingsModel.GetSentinelBoardArea()
	position1.board = make([]color.Color, boardMax)
	position1.checkBoard = make([]int, boardMax)
	position1.iteratorWithoutWall = position1.CreateBoardIteratorWithoutWall(readonlyGameSettingsModel)

	// 枠線
	for z := models.Point(0); z < models.Point(boardMax); z++ {
		position1.SetColor(z, color.Wall)
	}

	// 盤上
	var onPoint = func(z models.Point) {
		position1.SetColor(z, 0)
	}
	position1.iteratorWithoutWall(onPoint)

	position1.MovesNum = 0
	position1.KoZ = 0 // コウの指定がないので消します
}

// SetBoard - 盤面を設定します
func (position1 *Position) SetBoard(board []color.Color) {
	// TODO 消す
	// fmt.Print("[[")
	// for z := 0; z < SentinelBoardArea; z++ {
	// 	fmt.Printf("%d,", board[z])
	// 	position1.SetColor(Point(z), board[z])
	// }
	// fmt.Print("]]")
	position1.board = board
}

// ColorAt - 指定した交点の石の色
func (position1 *Position) ColorAt(z models.Point) color.Color {
	return position1.board[z]
}

// CheckAt - 指定した交点のチェック
func (position1 *Position) CheckAt(z models.Point) int {
	return position1.checkBoard[z]
}

// ColorAtXy - 指定した交点の石の色
func (position1 *Position) ColorAtXy(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, x int, y int) color.Color {
	return position1.board[(y+1)*readonlyGameSettingsModel.GetSentinelWidth()+x+1]
}

// IsEmpty - 指定の交点は空点か？
func (position1 *Position) IsEmpty(z models.Point) bool {
	return position1.board[z] == color.None
}

// SetColor - 盤データ。
func (position1 *Position) SetColor(z models.Point, color color.Color) {
	// TODO 消す
	// if z == Point(11) && color == Empty { // テスト
	// 	panic(fmt.Sprintf("z=%d color=%d SentinelWidth=%d", z, color, SentinelWidth))
	// }

	position1.board[z] = color
}

// GetZ4 - z（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (position1 *Position) GetZ4(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, z models.Point) int {
	if z == 0 {
		return 0
	}
	var y = int(z) / readonlyGameSettingsModel.GetSentinelWidth()
	var x = int(z) - y*readonlyGameSettingsModel.GetSentinelWidth()
	return x*100 + y
}

// GetZFromXy - x,y 形式の座標を、 z （配列のインデックス）へ変換します。
// x,y は壁を含まない領域での 0 から始まる座標です。 z は壁を含む盤上での座標です
func (position1 *Position) GetZFromXy(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, x int, y int) models.Point {
	return models.Point((y+1)*readonlyGameSettingsModel.GetSentinelWidth() + x + 1)
}

// GetEmptyZ - 空点の z （配列のインデックス）を返します。
func (position1 *Position) GetEmptyZ(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel) models.Point {
	var x, y int
	var z models.Point
	for {
		// ランダムに交点を選んで、空点を見つけるまで繰り返します。
		x = rand.Intn(readonlyGameSettingsModel.GetBoardSize()) // FIXME: 9 でいいの？ 9路盤？ → boardSize に変更
		y = rand.Intn(readonlyGameSettingsModel.GetBoardSize())
		z = position1.GetZFromXy(readonlyGameSettingsModel, x, y)
		if position1.IsEmpty(z) { // 空点
			break
		}
	}
	return z
}

// CountLiberty - 呼吸点を数えます。
// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (position1 *Position) CountLiberty(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, z models.Point, libertyArea *int, renArea *int) {
	*libertyArea = 0
	*renArea = 0

	// チェックボードの初期化
	var onPoint = func(z models.Point) {
		position1.checkBoard[z] = 0
	}
	position1.iteratorWithoutWall(onPoint)

	position1.countLibertySub(readonlyGameSettingsModel, z, position1.board[z], libertyArea, renArea)
}

// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (position1 *Position) countLibertySub(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, z models.Point, color color.Color, libertyArea *int, renArea *int) {
	position1.checkBoard[z] = 1
	*renArea++
	for i := 0; i < 4; i++ {
		var adjZ = z + readonlyGameSettingsModel.GetDirections4Array()[i]
		if position1.checkBoard[adjZ] != 0 {
			continue
		}
		if position1.IsEmpty(adjZ) { // 空点
			position1.checkBoard[adjZ] = 1
			*libertyArea++
		} else if position1.board[adjZ] == color {
			position1.countLibertySub(readonlyGameSettingsModel, adjZ, color, libertyArea, renArea) // 再帰
		}
	}
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (position1 *Position) TakeStone(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, z models.Point, color1 color.Color) {
	position1.board[z] = color.None // 石を消します

	for dir := 0; dir < 4; dir++ {
		var adjZ = z + readonlyGameSettingsModel.GetDirections4Array()[dir]

		if position1.board[adjZ] == color1 { // 再帰します
			position1.TakeStone(readonlyGameSettingsModel, adjZ, color1)
		}
	}
}

// IterateWithoutWall - 盤イテレーター
func (position1 *Position) IterateWithoutWall(onPoint func(models.Point)) {
	position1.iteratorWithoutWall(onPoint)
}

// UctChildrenSize - UCTの最大手数
func (position1 *Position) UctChildrenSize() int {
	return position1.uctChildrenSize
}

// CreateBoardIteratorWithoutWall - 盤の（壁を除く）全ての交点に順にアクセスする boardIterator 関数を生成します
func (position1 *Position) CreateBoardIteratorWithoutWall(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel) func(func(models.Point)) {

	var boardSize = readonlyGameSettingsModel.GetBoardSize()
	var boardIterator = func(onPoint func(models.Point)) {

		// UEC: 改造ポイント
		// 壁周りは後回しにしたい。
		if boardSize != 19 {
			// x,y は壁無しの盤での0から始まる座標、 z は壁有りの盤での0から始まる座標
			for y := 0; y < boardSize; y++ {
				for x := 0; x < boardSize; x++ {
					var z = position1.GetZFromXy(readonlyGameSettingsModel, x, y)
					onPoint(z)
				}
			}
		} else {
			// x,y は壁無しの盤での0から始まる座標、 z は壁有りの盤での0から始まる座標
			// for y := 0; y < boardSize; y++ {
			// 	for x := 0; x < boardSize; x++ {
			// 		var z = position1.GetZFromXy(readonlyGameSettingsModel, x, y)
			// 		onPoint(z)
			// 	}
			// }
			// FIMXE: 19 路盤：大会向けきふわらべ仕様
			// 改造： 数列
			numbers := []int{
				352, 353, 354, 355, 356, 357, 358, 359, 360, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298,
				351, 281, 282, 283, 284, 285, 286, 287, 288, 225, 226, 227, 228, 229, 230, 231, 232, 233, 299,
				350, 280, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 234, 300,
				349, 279, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 15, 235, 301,
				348, 278, 54, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 69, 16, 236, 302,
				347, 277, 53, 102, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 115, 70, 17, 237, 303,
				346, 276, 52, 101, 142, 175, 176, 177, 178, 179, 180, 181, 182, 153, 116, 71, 18, 238, 304,
				345, 275, 51, 100, 141, 174, 199, 200, 201, 202, 203, 204, 183, 154, 117, 72, 19, 239, 305,
				344, 274, 50, 99, 140, 173, 198, 215, 216, 217, 218, 205, 184, 155, 118, 73, 20, 240, 306,
				343, 273, 49, 98, 139, 172, 197, 214, 223, 224, 219, 206, 185, 156, 119, 74, 21, 241, 307,
				342, 272, 48, 97, 138, 171, 196, 213, 222, 221, 220, 207, 186, 157, 120, 75, 22, 242, 308,
				341, 271, 47, 96, 137, 170, 195, 212, 211, 210, 209, 208, 187, 158, 121, 76, 23, 243, 309,
				340, 270, 46, 95, 136, 169, 194, 193, 192, 191, 190, 189, 188, 159, 122, 77, 24, 244, 310,
				339, 269, 45, 94, 135, 168, 167, 166, 165, 164, 163, 162, 161, 160, 123, 78, 25, 245, 311,
				338, 268, 44, 93, 134, 133, 132, 131, 130, 129, 128, 127, 126, 125, 124, 79, 26, 246, 312,
				337, 267, 43, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 27, 247, 313,
				336, 266, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 248, 314,
				335, 265, 264, 263, 262, 261, 260, 259, 258, 257, 256, 255, 254, 253, 252, 251, 250, 249, 315,
				334, 333, 332, 331, 330, 329, 328, 327, 326, 325, 324, 323, 322, 321, 320, 319, 318, 317, 316,
			}
			for i := 0; i < boardSize*boardSize; i++ {
				var num = numbers[i]
				var y = num / 19
				var x = num % 19
				var z = position1.GetZFromXy(readonlyGameSettingsModel, x, y) // 壁を避けて計算
				onPoint(z)
			}
		}
	}

	return boardIterator
}

// PutStoneOnRecord - SelfPlay, RunGtpEngine から呼び出されます
func (position1 *Position) PutStoneOnRecord(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, z models.Point, color color.Color, recItem *gamerecord.GameRecordItem) {
	var err = position1.PutStone(readonlyGameSettingsModel, z, color)
	if err != 0 {
		logger.Console.Error("(PutStoneOnRecord) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	position1.Record[position1.MovesNum] = recItem
	position1.MovesNum++
}

// PutStone - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
//
// # Returns
// エラーコード
func (position1 *Position) PutStone(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, z models.Point, color1 color.Color) int {
	var around = [4]*ren.Ren{}   // 隣接する４つの交点
	var libertyArea int          // 呼吸点の数
	var renArea int              // 連の石の数
	var oppColor = color1.Flip() //相手(opponent)の石の色
	var space = 0                // 隣接している空点への向きの数
	var wall = 0                 // 隣接している壁への向きの数
	var myBreathFriend = 0       // 呼吸できる自分の石と隣接している向きの数
	var captureSum = 0           // アゲハマの数

	if z == models.Pass { // 投了なら、コウを消して関数を正常終了
		position1.KoZ = 0
		return 0
	}

	// 呼吸点を計算します
	for dir := 0; dir < 4; dir++ { // ４方向
		around[dir] = ren.NewRen(0, 0, 0) // 呼吸点の数, 連の石の数, 石の色

		var adjZ = z + readonlyGameSettingsModel.GetDirections4Array()[dir] // 隣の交点
		var adjColor = position1.ColorAt(adjZ)                              // 隣(adjacent)の交点の石の色
		if adjColor == color.None {                                         // 空点
			space++
			continue
		}
		if adjColor == color.Wall { // 壁
			wall++
			continue
		}
		position1.CountLiberty(readonlyGameSettingsModel, adjZ, &libertyArea, &renArea)
		around[dir].LibertyArea = libertyArea         // 呼吸点の数
		around[dir].StoneArea = renArea               // 連の意地の数
		around[dir].Color = adjColor                  // 石の色
		if adjColor == oppColor && libertyArea == 1 { // 相手の石で、呼吸点が１つで、その呼吸点に今石を置いたなら
			captureSum += renArea
		}
		if adjColor == color1 && 2 <= libertyArea { // 隣接する連が自分の石で、その石が呼吸点を２つ持ってるようなら
			myBreathFriend++
		}
	}

	// 石を置くと明らかに損なケース、また、ルール上石を置けないケースなら、石を置きません
	if captureSum == 0 && space == 0 && myBreathFriend == 0 {
		// 例えば黒番で 1 の箇所に打つのは損なので、石を置きません
		//
		//  ooo
		// ox1o
		//  oxo
		//   o
		return 1
	}
	if z == position1.KoZ { // コウには置けません
		return 2
	}
	if wall+myBreathFriend == 4 {
		// 例えば黒番で 1, 2 の箇所（眼）に打つのは損なので、石を置きません
		//
		// #########
		//  x2x  x1#
		//   x    x#
		//         #
		return 3
	}
	if !position1.IsEmpty(z) { // 空点以外には置けません
		return 4
	}

	position1.KoZ = 0 // コウを消します

	// 石を取り上げます
	for dir := 0; dir < 4; dir++ {
		var adjZ = z + readonlyGameSettingsModel.GetDirections4Array()[dir] // 隣接する交点
		var lib = around[dir].LibertyArea                                   // 隣接する連の呼吸点の数
		var adjColor = around[dir].Color                                    // 隣接する連の石の色

		if adjColor == oppColor && // 隣接する連が相手の石で（壁はここで除外されます）
			lib == 1 && // その呼吸点は１つで、そこに今石を置いた
			!position1.IsEmpty(adjZ) { // 石はまだあるなら（上と右の石がくっついている、といったことを除外）

			position1.TakeStone(readonlyGameSettingsModel, adjZ, oppColor)

			// もし取った石の数が１個なら、その石のある隣接した交点はコウ。また、図形上、コウは１個しか出現しません
			if around[dir].StoneArea == 1 {
				position1.KoZ = adjZ
			}
		}
	}

	position1.SetColor(z, color1)
	position1.CountLiberty(readonlyGameSettingsModel, z, &libertyArea, &renArea)

	return 0
}
