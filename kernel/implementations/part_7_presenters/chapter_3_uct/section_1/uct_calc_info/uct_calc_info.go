package uct_calc_info

import (
	"fmt"

	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_1/z_code"
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/model/gamesettingsmodel"
)

func CreatePrintingOfCalc(text_io1 i_text_io.ITextIO, readonlyGameSettingsModel *gamesettingsmodel.ObserverGameSettingsModel) *func(*position.Position, int, point.Point, float64, int) {
	// UCT計算中の表示
	var fn = func(position *position.Position, i int, z point.Point, rate float64, games int) {
		text_io1.LogInfo(fmt.Sprintf("(UCT Calculating...) %2d:z=%s,rate=%.4f,games=%3d\n", i, z_code.GetGtpZ(readonlyGameSettingsModel, position, z), rate, games))
	}

	return &fn
}

func CreatePrintingOfCalcFin(text_io1 i_text_io.ITextIO, readonlyGameSettingsModel *gamesettingsmodel.ObserverGameSettingsModel) *func(*position.Position, point.Point, float64, int, int, int) {
	// UCT計算後の表示
	var fn = func(position *position.Position, bestZ point.Point, rate float64, max int, allPlayouts int, nodeNum int) {
		text_io1.LogInfo(fmt.Sprintf("(UCT Calculated    ) bestZ=%s,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
			z_code.GetGtpZ(readonlyGameSettingsModel, position, bestZ), rate, max, allPlayouts, nodeNum))
	}

	return &fn
}
