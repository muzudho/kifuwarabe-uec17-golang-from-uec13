package mcts

import (
	// Entities

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models/color"

	// Use Cases
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/bademptytriangle"
)

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

var GettingOfWinnerOnDuringUCTPlayout *func(color.Color) int
var IsDislike *func(color.Color, models.Point) bool

// FlagTestPlayout - ？。
var FlagTestPlayout int

func InitPosition(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position1 *position.Position) {
	// 盤サイズが変わっていることもあるので、先に初期化します
	position1.InitPosition(readonlyGameSettingsModel)

	GettingOfWinnerOnDuringUCTPlayout = WrapGettingOfWinner(readonlyGameSettingsModel, position1)
	IsDislike = bademptytriangle.WrapIsDislike(readonlyGameSettingsModel, position1)

	AdjustParameters(readonlyGameSettingsModel, position1)
}
