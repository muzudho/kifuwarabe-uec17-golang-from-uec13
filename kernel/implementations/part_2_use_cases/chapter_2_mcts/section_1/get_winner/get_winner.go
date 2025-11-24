package get_winner

import (
	// Entities
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	komi_float "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/komi_float"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	gamesettingsmodel "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/model/gamesettingsmodel"
)

// WrapGettingOfWinner - 盤を束縛変数として与えます
func WrapGettingOfWinner(position *position.Position) *func(turnColor color.Color) int {
	// 「手番の勝ちなら1、引き分けなら0、手番の負けなら-1を返す関数（自分視点）」を作成します
	// * `turnColor` - 手番の石の色
	var getWinner = func(turnColor color.Color) int {
		return getWinner(position, turnColor)
	}

	return &getWinner
}

// 手番の勝ちなら1、引き分けなら0、手番の負けなら-1（自分視点）
// * `turnColor` - 手番の石の色
func getWinner(position *position.Position, turnColor color.Color) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	var onPoint = func(z point.Point) {
		var color2 = position.ColorAt(z)
		kind[color2]++
		if color2 == 0 {
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[position.ColorAt(z+gamesettingsmodel.Directions4Array[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}

	position.IterateWithoutWall(onPoint)

	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	var win = 0
	if 0 < komi_float.KomiFloat(score)-gamesettingsmodel.Komi {
		win = 1
	}
	if turnColor == 2 {
		win = -win
	} // gogo07

	return win
}
