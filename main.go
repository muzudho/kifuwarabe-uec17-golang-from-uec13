// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"
	"math/rand"
	"time"

	// 1. Entities
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	komi_float "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/komi_float"
	moves_num "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/moves_num"
	game_rule_settings "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_2_rule_settings/section_1"

	// 2. Use Cases
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"

	// 3. Controllers
	self_play "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_3_controllers/chapter_2_self_play/section_1/self_play"

	// 6. Gateways
	game_conf_toml "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_6_gateways/chapter_1_game_config/section_1/game_conf_toml"

	// 7. Presenters
	coding_obj "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_0_logger/section_1/coding_obj"
	text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_1_io/section_1"

	// Interfaces
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
)

func main() {
	flag.Parse()
	lessonVer := flag.Arg(0)

	// 乱数の種を設定
	rand.Seed(time.Now().UnixNano())

	// ログの書込み先設定
	coding_obj.GtpLog.SetPath("logs/gtp_print.log")
	coding_obj.ConsoleLog.SetPath(
		"logs/trace.log",
		"logs/debug.log",
		"logs/info.log",
		"logs/notice.log",
		"logs/warn.log",
		"logs/error.log",
		"logs/fatal.log",
		"logs/print.log")

	//coding_obj.Console.Trace("# Author: %s\n", game_rule_settings.Author)

	// 設定は囲碁GUIから与えられて上書きされる想定です。設定ファイルはデフォルト設定です
	var config = game_conf_toml.LoadGameConf("input/game_conf.toml", OnFatal)
	game_rule_settings.Komi = komi_float.KomiFloat(config.Komi())
	game_rule_settings.MaxMovesNum = moves_num.MovesNum(config.MaxMovesNum())
	game_rule_settings.SetBoardSize(config.BoardSize())
	var position = position.NewPosition()
	all_playouts.InitPosition(position)
	position.SetBoard(config.GetBoardArray())

	// ========================================
	// 思考エンジンの準備　＞　テキストＩＯ
	// ========================================

	var text_io1 i_text_io.ITextIO = text_io.NewTextIO()

	// ========================================
	// その他
	// ========================================

	if lessonVer == "SelfPlay" {
		self_play.SelfPlay(text_io1, position)
	} else {
		LoopGtp(text_io1, position) // GTP
	}
}

func OnFatal(errorMessage string) {
	coding_obj.Console.Fatal("%s", errorMessage)
}
