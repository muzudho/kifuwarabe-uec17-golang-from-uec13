package gamerecordpresenter

import (
	"fmt"

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamerecordusecase"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/textio"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"
)

func CreatePrintingOfCalc(text_io1 textio.ITextIO, readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel) *func(*position.Position, int, models.Point, float64, int) {
	// UCT計算中の表示
	var fn = func(position1 *position.Position, i int, z models.Point, rate float64, games int) {
		text_io1.LogInfo(fmt.Sprintf("(UCT Calculating...) %2d:z=%s,rate=%.4f,games=%3d\n", i, gamerecordusecase.GetGtpZ(readonlyGameSettingsModel, position1, z), rate, games))
	}

	return &fn
}

func CreatePrintingOfCalcFin(text_io1 textio.ITextIO, readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel) *func(*position.Position, models.Point, float64, int, int, int) {
	// UCT計算後の表示
	var fn = func(position1 *position.Position, bestZ models.Point, rate float64, max int, allPlayouts int, nodeNum int) {
		text_io1.LogInfo(fmt.Sprintf("(UCT Calculated    ) bestZ=%s,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
			gamerecordusecase.GetGtpZ(readonlyGameSettingsModel, position1, bestZ), rate, max, allPlayouts, nodeNum))
	}

	return &fn
}
