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
	e "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/entities"
	i_text_i_o "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_i_o/section_1/i_text_i_o"
	pl "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/play_algorithm"
	p "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/presenter"
)

// RunGtpEngine - レッスン９a
// GTP2NNGS に対応しているのでは？
func RunGtpEngine(text_i_o1 i_text_i_o.ITextIO, position *e.Position) {
	code.Console.Trace("# GoGo RunGtpEngine プログラム開始☆（＾～＾）\n")
	code.Console.Trace("# 何か標準入力しろだぜ☆（＾～＾）\n")

	// GUI から 囲碁エンジン へ入力があった、と考えてください
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var command = scanner.Text()
		code.Gtp.Log(command + "\n")
		code.ConsoleLog.Notice(command + "\n")

		var tokens = strings.Split(command, " ")
		switch tokens[0] {
		case "boardsize":
			// boardsize 19
			// 盤のサイズを変えます
			if 2 <= len(tokens) {
				var boardSize, err = strconv.Atoi(tokens[1])

				if err != nil {
					code.Console.Fatal(fmt.Sprintf("command=%s", command))
					panic(err)
				}

				e.SetBoardSize(boardSize)
				pl.InitPosition(position)

				text_i_o1.GoCommand("= \n\n")
			} else {
				text_i_o1.GoCommand(fmt.Sprintf("? unknown_command %s\n\n", command))
			}

		case "clear_board":
			pl.InitPosition(position)
			text_i_o1.GoCommand("= \n\n")

		case "quit":
			os.Exit(0)

		case "protocol_version":
			text_i_o1.GoCommand("= 2\n\n")

		case "name":
			text_i_o1.GoCommand("= Kifuwarabe UEC17 from UEC13\n\n")

		case "version":
			text_i_o1.GoCommand("= 0.0.2\n\n")

		case "list_commands":
			text_i_o1.GoCommand("= boardsize\nclear_board\nquit\nprotocol_version\nundo\n" +
				"name\nversion\nlist_commands\nkomi\ngenmove\nplay\n\n")

		case "komi":
			// komi 6.5
			if 2 <= len(tokens) {
				var komi, err = strconv.ParseFloat(tokens[1], 64)

				if err != nil {
					code.Console.Fatal(fmt.Sprintf("command=%s", command))
					panic(err)
				}

				e.Komi = e.KomiType(komi)
				text_i_o1.GoCommand(fmt.Sprintf("= %d\n\n", e.Komi))
			} else {
				text_i_o1.GoCommand(fmt.Sprintf("? unknown_command %s\n\n", command))
			}

			// TODO 消す text_i_o1.GoCommand("= 6.5\n\n")

		case "undo":
			// 未実装
			text_i_o1.GoCommand("= \n\n")

		case "genmove":
			// genmove black
			// genmove white
			var color e.Stone
			if 1 < len(tokens) && strings.ToLower(tokens[1][0:1]) == "w" {
				color = 2
			} else {
				color = 1
			}
			var z = PlayComputerMoveLesson09a(position, color)
			text_i_o1.GoCommand(fmt.Sprintf("= %s\n\n", p.GetGtpZ(position, z)))

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
				var color e.Stone
				if strings.ToLower(tokens[1][0:1]) == "w" {
					color = 2
				} else {
					color = 1
				}

				var z = p.GetZFromGtp(position, tokens[2])
				var recItem = new(e.RecordItem)
				recItem.Z = z
				recItem.Time = 0
				e.PutStoneOnRecord(position, z, color, recItem)
				p.PrintBoard(position, position.MovesNum)

				text_i_o1.GoCommand("= \n\n")
			}

		default:
			text_i_o1.GoCommand("? unknown_command\n\n")
		}
	}
}

// PlayComputerMoveLesson09a - コンピューター・プレイヤーの指し手。 SelfPlay, RunGtpEngine から呼び出されます。
func PlayComputerMoveLesson09a(
	position *e.Position,
	color e.Stone) e.Point {

	var st = time.Now()
	pl.AllPlayouts = 0

	var z, winRate = pl.GetBestZByUct(
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
		sec, float64(pl.AllPlayouts)/sec, position.GetZ4(z), winRate, position.MovesNum, color, pl.AllPlayouts)

	var recItem = new(e.RecordItem)
	recItem.Z = z
	recItem.Time = sec
	e.PutStoneOnRecord(position, z, color, recItem)
	p.PrintBoard(position, position.MovesNum)

	return z
}
