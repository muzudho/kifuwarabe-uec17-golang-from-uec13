package play_algorithm

import (
	e "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/entities"
	gd "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/game_domain"

	// Entities
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
)

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

var GettingOfWinnerOnDuringUCTPlayout *func(color.Color) int
var IsDislike *func(color.Color, point.Point) bool

// FlagTestPlayout - ？。
var FlagTestPlayout int

func InitPosition(position *e.Position) {
	// 盤サイズが変わっていることもあるので、先に初期化します
	position.InitPosition()

	GettingOfWinnerOnDuringUCTPlayout = WrapGettingOfWinner(position)
	IsDislike = gd.WrapIsDislike(position)
	AdjustParameters(position)
}
