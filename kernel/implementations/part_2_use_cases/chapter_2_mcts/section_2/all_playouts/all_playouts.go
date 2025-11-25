package all_playouts

import (
	// Entities
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models/color"

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

func InitPosition(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position1 *position.Position) {
	// 盤サイズが変わっていることもあるので、先に初期化します
	position1.InitPosition(readonlyGameSettingsModel)

	GettingOfWinnerOnDuringUCTPlayout = get_winner.WrapGettingOfWinner(readonlyGameSettingsModel, position1)
	IsDislike = bad_empty_triangle.WrapIsDislike(readonlyGameSettingsModel, position1)

	parameter_adjustment.AdjustParameters(readonlyGameSettingsModel, position1)
}
