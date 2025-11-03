package entities

import (
	"math/rand"
	"os"

	code "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/coding_obj"

	// Entities
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
	ren "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/ren"
	game_rule_settings "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_2_rule_settings/section_1/game_rule_settings"
)

// Position - 盤
type Position struct {
	// 盤
	board []color.Color
	// 呼吸点を数えるための一時盤
	checkBoard []int
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ point.Point
	// MovesNum - 手数
	MovesNum int
	// Record - 棋譜
	Record []*game_record_item.GameRecordItem
	// 二重ループ
	iteratorWithoutWall func(func(point.Point))
	// UCT計算中の子の数
	uctChildrenSize int
}

// TemporaryPosition - 盤をコピーするときの一時メモリーとして使います
type TemporaryPosition struct {
	// 盤
	Board []color.Color
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ point.Point
}

// CopyPosition - 盤データのコピー。
func (position *Position) CopyPosition() *TemporaryPosition {
	var temp = new(TemporaryPosition)
	temp.Board = make([]color.Color, game_rule_settings.SentinelBoardArea)
	copy(temp.Board[:], position.board[:])
	temp.KoZ = position.KoZ
	return temp
}

// ImportPosition - 盤データのコピー。
func (position *Position) ImportPosition(temp *TemporaryPosition) {
	copy(position.board[:], temp.Board[:])
	position.KoZ = temp.KoZ
}

// NewPosition - 空っぽの局面を生成します
// あとで InitPosition() を呼び出してください
func NewPosition() *Position {
	return new(Position)
}

// InitPosition - 局面の初期化。
func (position *Position) InitPosition() {
	position.Record = make([]*game_record_item.GameRecordItem, game_rule_settings.MaxMovesNum)
	position.uctChildrenSize = game_rule_settings.BoardArea + 1

	// サイズが変わっているケースに対応するため、配列の作り直し
	var boardMax = game_rule_settings.SentinelBoardArea
	position.board = make([]color.Color, boardMax)
	position.checkBoard = make([]int, boardMax)
	position.iteratorWithoutWall = CreateBoardIteratorWithoutWall(position)
	game_rule_settings.Directions4Array = [4]point.Point{1, -1, point.Point(game_rule_settings.SentinelWidth), point.Point(-game_rule_settings.SentinelWidth)}

	// 枠線
	for z := point.Point(0); z < point.Point(boardMax); z++ {
		position.SetColor(z, color.Wall)
	}

	// 盤上
	var onPoint = func(z point.Point) {
		position.SetColor(z, 0)
	}
	position.iteratorWithoutWall(onPoint)

	position.MovesNum = 0
	position.KoZ = 0 // コウの指定がないので消します
}

// SetBoard - 盤面を設定します
func (position *Position) SetBoard(board []color.Color) {
	// TODO 消す
	// fmt.Print("[[")
	// for z := 0; z < SentinelBoardArea; z++ {
	// 	fmt.Printf("%d,", board[z])
	// 	position.SetColor(Point(z), board[z])
	// }
	// fmt.Print("]]")
	position.board = board
}

// ColorAt - 指定した交点の石の色
func (position *Position) ColorAt(z point.Point) color.Color {
	return position.board[z]
}

// CheckAt - 指定した交点のチェック
func (position *Position) CheckAt(z point.Point) int {
	return position.checkBoard[z]
}

// ColorAtXy - 指定した交点の石の色
func (position *Position) ColorAtXy(x int, y int) color.Color {
	return position.board[(y+1)*game_rule_settings.SentinelWidth+x+1]
}

// IsEmpty - 指定の交点は空点か？
func (position *Position) IsEmpty(z point.Point) bool {
	return position.board[z] == color.None
}

// SetColor - 盤データ。
func (position *Position) SetColor(z point.Point, color color.Color) {
	// TODO 消す
	// if z == Point(11) && color == Empty { // テスト
	// 	panic(fmt.Sprintf("z=%d color=%d SentinelWidth=%d", z, color, SentinelWidth))
	// }

	position.board[z] = color
}

// GetZ4 - z（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (position *Position) GetZ4(z point.Point) int {
	if z == 0 {
		return 0
	}
	var y = int(z) / game_rule_settings.SentinelWidth
	var x = int(z) - y*game_rule_settings.SentinelWidth
	return x*100 + y
}

// GetZFromXy - x,y 形式の座標を、 z （配列のインデックス）へ変換します。
// x,y は壁を含まない領域での 0 から始まる座標です。 z は壁を含む盤上での座標です
func (position *Position) GetZFromXy(x int, y int) point.Point {
	return point.Point((y+1)*game_rule_settings.SentinelWidth + x + 1)
}

// GetEmptyZ - 空点の z （配列のインデックス）を返します。
func (position *Position) GetEmptyZ() point.Point {
	var x, y int
	var z point.Point
	for {
		// ランダムに交点を選んで、空点を見つけるまで繰り返します。
		x = rand.Intn(9)
		y = rand.Intn(9)
		z = position.GetZFromXy(x, y)
		if position.IsEmpty(z) { // 空点
			break
		}
	}
	return z
}

// CountLiberty - 呼吸点を数えます。
// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (position *Position) CountLiberty(z point.Point, libertyArea *int, renArea *int) {
	*libertyArea = 0
	*renArea = 0

	// チェックボードの初期化
	var onPoint = func(z point.Point) {
		position.checkBoard[z] = 0
	}
	position.iteratorWithoutWall(onPoint)

	position.countLibertySub(z, position.board[z], libertyArea, renArea)
}

// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (position *Position) countLibertySub(z point.Point, color color.Color, libertyArea *int, renArea *int) {
	position.checkBoard[z] = 1
	*renArea++
	for i := 0; i < 4; i++ {
		var adjZ = z + game_rule_settings.Directions4Array[i]
		if position.checkBoard[adjZ] != 0 {
			continue
		}
		if position.IsEmpty(adjZ) { // 空点
			position.checkBoard[adjZ] = 1
			*libertyArea++
		} else if position.board[adjZ] == color {
			position.countLibertySub(adjZ, color, libertyArea, renArea) // 再帰
		}
	}
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (position *Position) TakeStone(z point.Point, color1 color.Color) {
	position.board[z] = color.None // 石を消します

	for dir := 0; dir < 4; dir++ {
		var adjZ = z + game_rule_settings.Directions4Array[dir]

		if position.board[adjZ] == color1 { // 再帰します
			position.TakeStone(adjZ, color1)
		}
	}
}

// IterateWithoutWall - 盤イテレーター
func (position *Position) IterateWithoutWall(onPoint func(point.Point)) {
	position.iteratorWithoutWall(onPoint)
}

// UctChildrenSize - UCTの最大手数
func (position *Position) UctChildrenSize() int {
	return position.uctChildrenSize
}

// CreateBoardIteratorWithoutWall - 盤の（壁を除く）全ての交点に順にアクセスする boardIterator 関数を生成します
func CreateBoardIteratorWithoutWall(position *Position) func(func(point.Point)) {

	var boardSize = game_rule_settings.BoardSize
	var boardIterator = func(onPoint func(point.Point)) {

		// x,y は壁無しの盤での0から始まる座標、 z は壁有りの盤での0から始まる座標
		for y := 0; y < boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				var z = position.GetZFromXy(x, y)
				onPoint(z)
			}
		}
	}

	return boardIterator
}

// PutStoneOnRecord - SelfPlay, RunGtpEngine から呼び出されます
func PutStoneOnRecord(position *Position, z point.Point, color color.Color, recItem *game_record_item.GameRecordItem) {
	var err = PutStone(position, z, color)
	if err != 0 {
		code.Console.Error("(PutStoneOnRecord) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	position.Record[position.MovesNum] = recItem
	position.MovesNum++
}

// PutStone - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
//
// # Returns
// エラーコード
func PutStone(position *Position, z point.Point, color1 color.Color) int {
	var around = [4]*ren.Ren{}   // 隣接する４つの交点
	var libertyArea int          // 呼吸点の数
	var renArea int              // 連の石の数
	var oppColor = color1.Flip() //相手(opponent)の石の色
	var space = 0                // 隣接している空点への向きの数
	var wall = 0                 // 隣接している壁への向きの数
	var myBreathFriend = 0       // 呼吸できる自分の石と隣接している向きの数
	var captureSum = 0           // アゲハマの数

	if z == point.Pass { // 投了なら、コウを消して関数を正常終了
		position.KoZ = 0
		return 0
	}

	// 呼吸点を計算します
	for dir := 0; dir < 4; dir++ { // ４方向
		around[dir] = ren.NewRen(0, 0, 0) // 呼吸点の数, 連の石の数, 石の色

		var adjZ = z + game_rule_settings.Directions4Array[dir] // 隣の交点
		var adjColor = position.ColorAt(adjZ)                   // 隣(adjacent)の交点の石の色
		if adjColor == color.None {                             // 空点
			space++
			continue
		}
		if adjColor == color.Wall { // 壁
			wall++
			continue
		}
		position.CountLiberty(adjZ, &libertyArea, &renArea)
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
	if z == position.KoZ { // コウには置けません
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
	if !position.IsEmpty(z) { // 空点以外には置けません
		return 4
	}

	position.KoZ = 0 // コウを消します

	// 石を取り上げます
	for dir := 0; dir < 4; dir++ {
		var adjZ = z + game_rule_settings.Directions4Array[dir] // 隣接する交点
		var lib = around[dir].LibertyArea                       // 隣接する連の呼吸点の数
		var adjColor = around[dir].Color                        // 隣接する連の石の色

		if adjColor == oppColor && // 隣接する連が相手の石で（壁はここで除外されます）
			lib == 1 && // その呼吸点は１つで、そこに今石を置いた
			!position.IsEmpty(adjZ) { // 石はまだあるなら（上と右の石がくっついている、といったことを除外）

			position.TakeStone(adjZ, oppColor)

			// もし取った石の数が１個なら、その石のある隣接した交点はコウ。また、図形上、コウは１個しか出現しません
			if around[dir].StoneArea == 1 {
				position.KoZ = adjZ
			}
		}
	}

	position.SetColor(z, color1)
	position.CountLiberty(z, &libertyArea, &renArea)

	return 0
}
