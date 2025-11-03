package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	code "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/coding_obj"
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
	p "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/presenter"

	// Entities
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	komi_float "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/komi_float"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
	game_rule_settings "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_2_rule_settings/section_1/game_rule_settings"
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_3_position/section_1/position"
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_4/uct"
)

// LoopGtp - レッスン９a
// GTP2NNGS に対応しているのでは？
func LoopGtp(text_io1 i_text_io.ITextIO, position *position.Position) {
	code.Console.Trace("# GoGo RunGtpEngine プログラム開始☆（＾～＾）\n")
	code.Console.Trace("# 何か標準入力しろだぜ☆（＾～＾）\n")

	// GUI から 囲碁エンジン へ入力があった、と考えてください
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var command = scanner.Text()
		text_io1.ReceivedCommand(command)

		var tokens = strings.Split(command, " ")
		switch tokens[0] {

		// ========================================
		// GTP 対応　＞　大会参加最低限
		// ========================================

		case "list_commands":
			// 最初の１個は頭に "= " を付ける必要があってめんどくさいので先に出力
			text_io1.SendCommand("= list_commands\n")

			items := []string{
				// 終了コマンド
				"quit",
				// ハンドシェイク
				"protocol_version", "name", "version",
				// 対局設定
				"boardsize", "komi",
				// 対局
				"clear_board", "play", "undo", "genmove"}
			for _, item := range items {
				text_io1.SendCommand(fmt.Sprintf("%s\n", item))
			}

			text_io1.SendCommand("\n")

		// ========================================
		// GTP 対応　＞　大会参加最低限　＞　終了コマンド
		// ========================================

		case "quit":
			os.Exit(0)

		// ========================================
		// GTP 対応　＞　大会参加最低限　＞　ハンドシェイク
		// ========================================

		case "protocol_version":
			text_io1.SendCommand("= 2\n\n")

		case "name":
			text_io1.SendCommand("= Kifuwarabe UEC17 from UEC13\n\n")

		case "version":
			text_io1.SendCommand("= 0.0.2\n\n")

		// ========================================
		// GTP 対応　＞　大会参加最低限　＞　対局設定
		// ========================================

		case "boardsize":
			// boardsize 19
			// 盤のサイズを変えます
			if 2 <= len(tokens) {
				var boardSize, err = strconv.Atoi(tokens[1])

				if err != nil {
					code.Console.Fatal(fmt.Sprintf("command=%s", command))
					panic(err)
				}

				game_rule_settings.SetBoardSize(boardSize)
				all_playouts.InitPosition(position)

				text_io1.SendCommand("= \n\n")
			} else {
				text_io1.SendCommand(fmt.Sprintf("? unknown_command %s\n\n", command))
			}

		case "komi":
			// komi 6.5
			if 2 <= len(tokens) {
				var komi, err = strconv.ParseFloat(tokens[1], 64)

				if err != nil {
					code.Console.Fatal(fmt.Sprintf("command=%s", command))
					panic(err)
				}

				game_rule_settings.Komi = komi_float.KomiFloat(komi)
				text_io1.SendCommand(fmt.Sprintf("= %f\n\n", game_rule_settings.Komi))
			} else {
				text_io1.SendCommand(fmt.Sprintf("? unknown_command %s\n\n", command))
			}

			// TODO 消す text_io1.SendCommand("= 6.5\n\n")

		// ========================================
		// GTP 対応　＞　大会参加最低限　＞　対局
		// ========================================

		case "clear_board":
			all_playouts.InitPosition(position)
			text_io1.SendCommand("= \n\n")

		case "play":
			// play black A3
			// play white D4
			// play black D5
			// play white E5
			// play black E4
			// play white D6
			// play black F5
			// play white C5
			// play black PASS
			// play white PASS
			if 2 < len(tokens) {
				var color color.Color
				if strings.ToLower(tokens[1][0:1]) == "w" {
					color = 2
				} else {
					color = 1
				}

				var z = p.GetZFromGtp(position, tokens[2])
				var recItem = new(game_record_item.GameRecordItem)
				recItem.Z = z
				recItem.Time = 0
				position.PutStoneOnRecord(z, color, recItem)
				p.PrintBoard(position, position.MovesNum)

				text_io1.SendCommand("= \n\n")
			}

		case "undo":
			// 未実装
			text_io1.SendCommand("= \n\n")

		case "genmove":
			// genmove black
			// genmove white
			var color color.Color
			if 1 < len(tokens) && strings.ToLower(tokens[1][0:1]) == "w" {
				color = 2
			} else {
				color = 1
			}
			var z = PlayComputerMoveLesson09a(position, color)
			text_io1.SendCommand(fmt.Sprintf("= %s\n\n", p.GetGtpZ(position, z)))

		default:
			text_io1.SendCommand("? unknown_command\n\n")
		}
	}
}

// PlayComputerMoveLesson09a - コンピューター・プレイヤーの指し手。 SelfPlay, RunGtpEngine から呼び出されます。
func PlayComputerMoveLesson09a(
	position *position.Position,
	color color.Color) point.Point {

	var st = time.Now()
	all_playouts.AllPlayouts = 0

	var z, winRate = uct.GetBestZByUct(
		position,
		color,
		createPrintingOfCalc(),
		createPrintingOfCalcFin())

	if 1 < position.MovesNum && // 初手ではないとして
		position.Record[position.MovesNum-1].GetZ() == 0 && // １つ前の手がパスで
		0.95 <= math.Abs(winRate) { // 95%以上の確率で勝ちか負けなら
		// こちらもパスします
		return 0
	}

	var sec = time.Since(st).Seconds()
	code.Console.Info("%.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(all_playouts.AllPlayouts)/sec, position.GetZ4(z), winRate, position.MovesNum, color, all_playouts.AllPlayouts)

	var recItem = new(game_record_item.GameRecordItem)
	recItem.Z = z
	recItem.Time = sec
	position.PutStoneOnRecord(z, color, recItem)
	p.PrintBoard(position, position.MovesNum)

	return z
}
