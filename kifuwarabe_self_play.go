package main

import (
	"time"

	// Entities
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_3_position/section_1/position"
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_4/uct"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_0_logger/section_1/coding_obj"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_1/z_code"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_2/sgf"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_3/board_view"
)

// SelfPlay - コンピューター同士の対局。
func SelfPlay(position *position.Position) {
	coding_obj.Console.Trace("# GoGo SelfPlay 自己対局開始☆（＾～＾）\n")

	var color = color.Black

	for {
		var z = GetComputerMoveDuringSelfPlay(position, color)

		var recItem = new(game_record_item.GameRecordItem)
		recItem.Z = z
		position.PutStoneOnRecord(z, color, recItem)

		coding_obj.Console.Print("z=%s,color=%d", z_code.GetGtpZ(position, z), color) // テスト
		// p.PrintCheckBoard(position)                                        // テスト
		board_view.PrintBoard(position, position.MovesNum)

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

	sgf.PrintSgf(position, position.MovesNum, position.Record)
}

// GetComputerMoveDuringSelfPlay - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveDuringSelfPlay(position *position.Position, color color.Color) point.Point {

	var start = time.Now()
	all_playouts.AllPlayouts = 0

	var z, winRate = uct.GetBestZByUct(
		position,
		color,
		createPrintingOfCalc(),
		createPrintingOfCalcFin())

	var sec = time.Since(start).Seconds()
	coding_obj.Console.Info("(GetComputerMoveDuringSelfPlay) %.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(all_playouts.AllPlayouts)/sec, position.GetZ4(z), winRate, position.MovesNum, color, all_playouts.AllPlayouts)
	return z
}
