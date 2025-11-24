package self_play

import (
	"fmt"
	"time"

	// Entities
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
gamesettings
	// User Cases
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"
	uct "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_4/uct"

	// Presenters
	coding_obj "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_0_logger/section_1/coding_obj"
	z_code "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_1/z_code"
	sgf "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_2/sgf"
	board_view "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_3/board_view"
	uct_calc_info "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_3_uct/section_1/uct_calc_info"

	// Interfaces
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
)

// SelfPlay - コンピューター同士の対局。
func SelfPlay(text_io1 i_text_io.ITextIO, readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position *position.Position) {
	coding_obj.Console.Trace("# GoGo SelfPlay 自己対局開始☆（＾～＾）\n")

	var color = color.Black

	for {
		var z = GetComputerMoveDuringSelfPlay(text_io1, readonlyGameSettingsModel, position, color)

		var recItem = new(game_record_item.GameRecordItem)
		recItem.Z = z
		position.PutStoneOnRecord(readonlyGameSettingsModel, z, color, recItem)

		coding_obj.Console.Print("z=%s,color=%d", z_code.GetGtpZ(readonlyGameSettingsModel, position, z), color) // テスト

		// p.PrintCheckBoard(readonlyGameSettingsModel, position)                                        // テスト
		board_view.PrintBoard(readonlyGameSettingsModel, position, position.MovesNum)

		// パスで２手目以降で棋譜の１つ前（相手）もパスなら終了します。
		if z == point.Pass && 1 < position.MovesNum && position.Record[position.MovesNum-2].GetZ() == point.Pass {
			break
		}
		// 自己対局は400手で終了します。
		if 400 < position.MovesNum {
			break
		} // too long
		color = color.Flip()
	}

	sgf.PrintSgf(readonlyGameSettingsModel, position, position.MovesNum, position.Record)
}

// GetComputerMoveDuringSelfPlay - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveDuringSelfPlay(text_io1 i_text_io.ITextIO, readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position *position.Position, color color.Color) point.Point {

	var start = time.Now()
	all_playouts.AllPlayouts = 0

	var z, winRate = uct.GetBestZByUct(
		readonlyGameSettingsModel,
		position,
		color,
		uct_calc_info.CreatePrintingOfCalc(text_io1, readonlyGameSettingsModel),
		uct_calc_info.CreatePrintingOfCalcFin(text_io1, readonlyGameSettingsModel))

	var sec = time.Since(start).Seconds()
	text_io1.LogInfo(fmt.Sprintf("(GetComputerMoveDuringSelfPlay) %.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(all_playouts.AllPlayouts)/sec, position.GetZ4(readonlyGameSettingsModel, z), winRate, position.MovesNum, color, all_playouts.AllPlayouts))
	return z
}
