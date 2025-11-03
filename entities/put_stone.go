package entities

import (
	"os"

	code "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/coding_obj"

	// Entities
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
	ren "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/ren"
)

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

		var adjZ = z + Directions4Array[dir]  // 隣の交点
		var adjColor = position.ColorAt(adjZ) // 隣(adjacent)の交点の石の色
		if adjColor == color.None {           // 空点
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
		var adjZ = z + Directions4Array[dir] // 隣接する交点
		var lib = around[dir].LibertyArea    // 隣接する連の呼吸点の数
		var adjColor = around[dir].Color     // 隣接する連の石の色

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
