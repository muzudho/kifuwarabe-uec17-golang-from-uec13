// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/textio"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mcts"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mctsimpl"
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
	var position1 = position.NewPosition()
	mcts.InitPosition(readonlyGameSettingsModel, position1)
	position1.SetBoard(gamesettings.GetBoardArray(&dto1))

	// ========================================
	// 思考エンジンの準備　＞　テキストＩＯ
	// ========================================

	var text_io1 textio.ITextIO = textio.NewTextIO()

	// ========================================
	// その他
	// ========================================

	if lessonVer == "SelfPlay" {
		mctsimpl.SelfPlay(text_io1, readonlyGameSettingsModel, position1)
	} else {
		LoopGtp(text_io1, &dto1, position1) // GTP
	}
}

func OnFatal(errorMessage string) {
	logger.Console.Fatal("%s", errorMessage)
}
