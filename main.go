// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_6_gateways/chapter_1_game_config/section_1/game_conf_toml"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_0_logger/section_1/coding_obj"
	text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_1_io/section_1"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_1/z_code"
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"

	// Entity
	komi_float "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/komi_float"
	moves_num "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/moves_num"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	game_rule_settings "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_2_rule_settings/section_1/game_rule_settings"
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_3_position/section_1/position"
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"
)

func main() {
	flag.Parse()
	lessonVer := flag.Arg(0)

	// 乱数の種を設定
	rand.Seed(time.Now().UnixNano())

	// ログの書込み先設定
	coding_obj.GtpLog.SetPath("output/gtp_print.log")
	coding_obj.ConsoleLog.SetPath(
		"output/trace.log",
		"output/debug.log",
		"output/info.log",
		"output/notice.log",
		"output/warn.log",
		"output/error.log",
		"output/fatal.log",
		"output/print.log")

	coding_obj.Console.Trace("# Author: %s\n", game_rule_settings.Author)

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
		SelfPlay(position)
	} else {
		LoopGtp(text_io1, position) // GTP
	}
}

func OnFatal(errorMessage string) {
	coding_obj.Console.Fatal(errorMessage)
}

func createPrintingOfCalc() *func(*position.Position, int, point.Point, float64, int) {
	// UCT計算中の表示
	var fn = func(position *position.Position, i int, z point.Point, rate float64, games int) {
		coding_obj.Console.Info("(UCT Calculating...) %2d:z=%s,rate=%.4f,games=%3d\n", i, z_code.GetGtpZ(position, z), rate, games)
	}

	return &fn
}

func createPrintingOfCalcFin() *func(*position.Position, point.Point, float64, int, int, int) {
	// UCT計算後の表示
	var fn = func(position *position.Position, bestZ point.Point, rate float64, max int, allPlayouts int, nodeNum int) {
		coding_obj.Console.Info("(UCT Calculated    ) bestZ=%s,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
			z_code.GetGtpZ(position, bestZ), rate, max, allPlayouts, nodeNum)

	}

	return &fn
}
