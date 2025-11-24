package game_domain

// Empty triangle (アキ三角)
//
// x.
// xx

import (
	// Entities
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	direction_4 "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/direction_4"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
)

// WrapIsDislike - 盤を束縛変数として与えます
func WrapIsDislike(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position1 *position.Position) *func(color.Color, point.Point) bool {
	// 「手番の勝ちなら1、引き分けなら0、手番の負けなら-1を返す関数（自分視点）」を作成します
	// * `color` - 石の色
	var isDislike = func(color color.Color, z point.Point) bool {
		// 座標取得
		// 432
		// 5S1
		// 678
		var eastZ = z + readonlyGameSettingsModel.GetDirections4Array()[direction_4.East]
		var northEastZ = z + readonlyGameSettingsModel.GetDirections4Array()[direction_4.North] + 1
		var northZ = z + readonlyGameSettingsModel.GetDirections4Array()[direction_4.North]
		var northWestZ = z + readonlyGameSettingsModel.GetDirections4Array()[direction_4.North] - 1
		var westZ = z + readonlyGameSettingsModel.GetDirections4Array()[direction_4.West]
		var southWestZ = z + readonlyGameSettingsModel.GetDirections4Array()[direction_4.South] - 1
		var southZ = z + readonlyGameSettingsModel.GetDirections4Array()[direction_4.South]
		var southEastZ = z + readonlyGameSettingsModel.GetDirections4Array()[direction_4.South] + 1

		// 東北
		// **
		// S*
		if isEmptyTriangle(position1, color, [3]point.Point{eastZ, northEastZ, northZ}) {
			return true
		}

		// 西北
		// **
		// *S
		if isEmptyTriangle(position1, color, [3]point.Point{northZ, northWestZ, westZ}) {
			return true
		}

		// 西南
		// *S
		// **
		if isEmptyTriangle(position1, color, [3]point.Point{westZ, southWestZ, southZ}) {
			return true
		}

		// 東南
		// S*
		// **
		if isEmptyTriangle(position1, color, [3]point.Point{southZ, southEastZ, eastZ}) {
			return true
		}

		return false
	}

	return &isDislike
}

func isEmptyTriangle(position1 *position.Position, myColor color.Color, points [3]point.Point) bool {
	var myColorNum = 0
	var emptyNum = 0

	for _, z := range points {
		var color1 = position1.ColorAt(z)
		switch color1 {
		case myColor:
			myColorNum++
		case color.None:
			emptyNum++
		}
	}

	return myColorNum == 2 && emptyNum == 1
}
