package play_computer_move_lesson_09_a

import (
	"math"
	"time"

	// Entities
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/model/gamesettingsmodel"

	// Use Cases
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"
	uct "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_4/uct"

	// Presenters

	uct_calc_info "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_3_uct/section_1/uct_calc_info"

	// Interfaces
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
)

// PlayComputerMoveLesson09a - コンピューター・プレイヤーの指し手。 SelfPlay, RunGtpEngine から呼び出されます。
func PlayComputerMoveLesson09a(
	text_io1 i_text_io.ITextIO,
	readonlyGameSettingsModel *gamesettingsmodel.ReadonlyGameSettingsModel,
	position1 *position.Position,
	color1 color.Color) point.Point {

	var st1 = time.Now()
	all_playouts.AllPlayouts = 0

	var z1, winRate1 = uct.GetBestZByUct(
		readonlyGameSettingsModel,
		position1,
		color1,
		uct_calc_info.CreatePrintingOfCalc(text_io1, readonlyGameSettingsModel),
		uct_calc_info.CreatePrintingOfCalcFin(text_io1, readonlyGameSettingsModel))

	if 1 < position1.MovesNum && // 初手ではないとして
		position1.Record[position1.MovesNum-1].GetZ() == 0 && // １つ前の手がパスで
		0.95 <= math.Abs(winRate1) { // 95%以上の確率で勝ちか負けなら
		// こちらもパスします
		return 0
	}

	var sec1 = time.Since(st1).Seconds()
	// FIXME: text_io1.LogInfo(fmt.Sprintf("%.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,movesNum=%d,color=%d,playouts=%d\n",
	//		sec1, float64(all_playouts.AllPlayouts)/sec1, position1.GetZ4(z1), winRate1, position1.MovesNum, color1, all_playouts.AllPlayouts))

	var recItem1 = new(game_record_item.GameRecordItem)
	recItem1.Z = z1
	recItem1.Time = sec1
	position1.PutStoneOnRecord(readonlyGameSettingsModel, z1, color1, recItem1)

	// FIXME: board_view.PrintBoard(readonlyGameSettingsModel, position1, position1.MovesNum)

	return z1
}
