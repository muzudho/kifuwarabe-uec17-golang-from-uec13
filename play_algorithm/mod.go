package play_algorithm

import (
	// Entities
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_3_position/section_1/position"

	// Use Cases
	bad_empty_triangle "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_1_game_domain/section_1/bad_empty_triangle"
)

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

var GettingOfWinnerOnDuringUCTPlayout *func(color.Color) int
var IsDislike *func(color.Color, point.Point) bool

// FlagTestPlayout - ？。
var FlagTestPlayout int

func InitPosition(position *position.Position) {
	// 盤サイズが変わっていることもあるので、先に初期化します
	position.InitPosition()

	GettingOfWinnerOnDuringUCTPlayout = WrapGettingOfWinner(position)
	IsDislike = bad_empty_triangle.WrapIsDislike(position)
	AdjustParameters(position)
}
