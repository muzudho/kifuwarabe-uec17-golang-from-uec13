package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	// 1 Entities
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_3_controllers/chapter_1_computer_player/section_1/play_computer_move_lesson_09_a"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"

	// 2 Use Cases
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"

	// 7 Presenters
	z_code "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_1/z_code"
	board_view "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_3/board_view"
	coding_obj "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"

	// Interfaces
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
)

// LoopGtp - レッスン９a
// GTP2NNGS に対応しているのでは？
func LoopGtp(text_io1 i_text_io.ITextIO, gameSettingsDto1 *gamesettings.GameSettingsFile, position *position.Position) {
	//coding_obj.Console.Trace("# きふわらべ UEC17 golang from UEC13 プログラム開始☆（＾～＾）\n")
	//coding_obj.Console.Trace("# 何か標準入力しろだぜ☆（＾～＾）\n")

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
			// ```shell
			// quit
			// ```
			os.Exit(0)

		// ========================================
		// GTP 対応　＞　大会参加最低限　＞　ハンドシェイク
		// ========================================

		case "protocol_version":
			// ```shell
			// protocol_version
			// ```
			text_io1.SendCommand("= 2\n\n")

		case "name":
			// ```shell
			// name
			// ```
			text_io1.SendCommand("= Kifuwarabe UEC17 from UEC13\n\n")

		case "version":
			// ```shell
			// version
			// ```
			text_io1.SendCommand("= 0.0.2\n\n")

		// ========================================
		// GTP 対応　＞　大会参加最低限　＞　対局設定
		// ========================================

		case "boardsize":
			// ```shell
			// boardsize 19
			// ```
			// 盤のサイズを変えます
			if 2 <= len(tokens) {
				var boardSize, err = strconv.Atoi(tokens[1])

				if err != nil {
					coding_obj.Console.Fatal("command=%s", command)
					panic(err)
				}

				gameSettingsDto1.Game.BoardSize = int8(boardSize)
				var readonlyGameSettingsModel = gamesettings.NewReadonlyGameSettingsModel(gameSettingsDto1.Game.GetBoardSize(), gameSettingsDto1.Game.GetKomi(), gameSettingsDto1.Game.GetMaxMoves())
				all_playouts.InitPosition(readonlyGameSettingsModel, position)

				text_io1.SendCommand("= \n\n")
			} else {
				text_io1.SendCommand(fmt.Sprintf("? unknown_command %s\n\n", command))
			}

		case "komi":
			// ```shell
			// komi 6.5
			// ```
			if 2 <= len(tokens) {
				var komi, err = strconv.ParseFloat(tokens[1], 64)

				if err != nil {
					coding_obj.Console.Fatal("command=%s", command)
					panic(err)
				}

				gameSettingsDto1.Game.Komi = float32(komi)
				text_io1.SendCommand(fmt.Sprintf("= %g\n\n", gameSettingsDto1.Game.Komi))
			} else {
				text_io1.SendCommand(fmt.Sprintf("? unknown_command %s\n\n", command))
			}

			// TODO 消す text_io1.SendCommand("= 6.5\n\n")

		// ========================================
		// GTP 対応　＞　大会参加最低限　＞　対局
		// ========================================

		case "clear_board":
			// ```shell
			// clear_board
			// ```
			var readonlyGameSettingsModel = gamesettings.NewReadonlyGameSettingsModel(gameSettingsDto1.Game.GetBoardSize(), gameSettingsDto1.Game.GetKomi(), gameSettingsDto1.Game.GetMaxMoves())
			all_playouts.InitPosition(readonlyGameSettingsModel, position)
			text_io1.SendCommand("= \n\n")

		case "play":
			// ```shell
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
			// ```
			if 2 < len(tokens) {
				var color color.Color
				if strings.ToLower(tokens[1][0:1]) == "w" {
					color = 2
				} else {
					color = 1
				}

				var readonlyGameSettingsModel = gamesettings.NewReadonlyGameSettingsModel(gameSettingsDto1.Game.GetBoardSize(), gameSettingsDto1.Game.GetKomi(), gameSettingsDto1.Game.GetMaxMoves())
				var z = z_code.GetZFromGtp(readonlyGameSettingsModel, position, tokens[2])
				var recItem = new(game_record_item.GameRecordItem)
				recItem.Z = z
				recItem.Time = 0
				position.PutStoneOnRecord(readonlyGameSettingsModel, z, color, recItem)

				text_io1.SendCommand("= \n\n")
			}

		case "undo":
			// 未実装
			text_io1.SendCommand("= \n\n")

		case "genmove":
			// ```shell
			// genmove black
			// genmove white
			// ```
			var color1 color.Color
			if 1 < len(tokens) && strings.ToLower(tokens[1][0:1]) == "w" {
				color1 = 2
			} else {
				color1 = 1
			}

			var readonlyGameSettingsModel = gamesettings.NewReadonlyGameSettingsModel(gameSettingsDto1.Game.GetBoardSize(), gameSettingsDto1.Game.GetKomi(), gameSettingsDto1.Game.GetMaxMoves())
			var z = play_computer_move_lesson_09_a.PlayComputerMoveLesson09a(text_io1, readonlyGameSettingsModel, position, color1)
			text_io1.SendCommand(fmt.Sprintf("= %s\n\n", z_code.GetGtpZ(readonlyGameSettingsModel, position, z)))

		// ========================================
		// 独自実装
		// ========================================

		case "-board":
			// ```shell
			// -board
			// ```
			text_io1.SendCommand("= \n")
			var readonlyGameSettingsModel = gamesettings.NewReadonlyGameSettingsModel(gameSettingsDto1.Game.GetBoardSize(), gameSettingsDto1.Game.GetKomi(), gameSettingsDto1.Game.GetMaxMoves())
			board_view.PrintBoard(readonlyGameSettingsModel, position, position.MovesNum)
			text_io1.SendCommand("\n\n")

		default:
			text_io1.SendCommand("? unknown_command\n\n")
		}
	}
}
