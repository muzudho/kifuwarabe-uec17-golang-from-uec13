package all_playouts

import (
	// Entities
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/model/gamesettingsmodel"

	// Use Cases
	bad_empty_triangle "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_1_game_domain/section_1/bad_empty_triangle"
	get_winner "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/get_winner"
	parameter_adjustment "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/parameter_adjustment"
)

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

var GettingOfWinnerOnDuringUCTPlayout *func(color.Color) int
var IsDislike *func(color.Color, point.Point) bool

// FlagTestPlayout - ？。
var FlagTestPlayout int

func InitPosition(observerGameSettingsModel *gamesettingsmodel.ObserverGameSettingsModel, position *position.Position) {
	// 盤サイズが変わっていることもあるので、先に初期化します
	position.InitPosition(observerGameSettingsModel)

	GettingOfWinnerOnDuringUCTPlayout = get_winner.WrapGettingOfWinner(observerGameSettingsModel, position)
	IsDislike = bad_empty_triangle.WrapIsDislike(position)

	parameter_adjustment.AdjustParameters(observerGameSettingsModel, position)
}
