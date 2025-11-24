// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"
	"math/rand"
	"time"

	// 1. Entities
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"

	// 2. Use Cases
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"

	// 3. Controllers
	self_play "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_3_controllers/chapter_2_self_play/section_1/self_play"

	// 6. Gateways
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"

	// 7. Presenters
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/textio"

	// Interfaces
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
)

func main() {
	flag.Parse()
	lessonVer := flag.Arg(0)

	// 乱数の種を設定
	rand.Seed(time.Now().UnixNano())

	// ログの書込み先設定
	logger.GtpLog.SetPath("logs/gtp_print.log")
	logger.ConsoleLog.SetPath(
		"logs/trace.log",
		"logs/debug.log",
		"logs/info.log",
		"logs/notice.log",
		"logs/warn.log",
		"logs/error.log",
		"logs/fatal.log",
		"logs/print.log")

	//logger.Console.Trace("# Author: %s\n", gamesettings.Author)

	// 設定は囲碁GUIから与えられて上書きされる想定です。設定ファイルはデフォルト設定です
	var dto1 = gamesettings.LoadGameSettings("game_settings.toml", OnFatal)
	var readonlyGameSettingsModel = gamesettings.NewReadonlyGameSettingsModel(dto1.Game.GetBoardSize(), dto1.Game.GetKomi(), dto1.Game.GetMaxMoves())
	var position = position.NewPosition()
	all_playouts.InitPosition(readonlyGameSettingsModel, position)
	position.SetBoard(gamesettings.GetBoardArray(&dto1))

	// ========================================
	// 思考エンジンの準備　＞　テキストＩＯ
	// ========================================

	var text_io1 i_text_io.ITextIO = textio.NewTextIO()

	// ========================================
	// その他
	// ========================================

	if lessonVer == "SelfPlay" {
		self_play.SelfPlay(text_io1, readonlyGameSettingsModel, position)
	} else {
		LoopGtp(text_io1, &dto1, position) // GTP
	}
}

func OnFatal(errorMessage string) {
	logger.Console.Fatal("%s", errorMessage)
}
