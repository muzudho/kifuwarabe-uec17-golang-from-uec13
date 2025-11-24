package uct

import (
	"math"
	"math/rand"
	"os"

	// Entities
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	child "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/child"
	parameter_adjustment "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/parameter_adjustment"
	uct_struct "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/uct_struct"
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"
	node_struct "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_3/node_struct"
	playout "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_3/playout"
	coding_obj "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_0_logger/section_1/coding_obj"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/model/gamesettingsmodel"
)

// GetBestZByUct - Lesson08,09,09aで使用。 一番良いUCTである着手を選びます。 GetComputerMoveDuringSelfPlay などから呼び出されます。
//
// # Return
// (bestZ int, winRate float64)
func GetBestZByUct(
	observerGameSettingsModel *gamesettingsmodel.ObserverGameSettingsModel,
	position *position.Position,
	color color.Color,
	print_calc *func(*position.Position, int, point.Point, float64, int),
	print_calc_fin *func(*position.Position, point.Point, float64, int, int, int)) (point.Point, float64) {

	// UCT計算フェーズ
	node_struct.NodeNum = 0 // カウンターリセット
	var next = node_struct.CreateNode(position)
	var uctLoopCount = parameter_adjustment.UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		// 一時記憶
		var copiedPosition = position.CopyPosition(observerGameSettingsModel)

		SearchUct(observerGameSettingsModel, position, color, next)

		// 復元
		position.ImportPosition(copiedPosition)
	}

	// ベスト値検索フェーズ
	var bestI = -1
	var max = -999
	var pN = &node_struct.Nodes[next]
	for i := 0; i < pN.ChildNum; i++ {
		var c = &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		// FIXME: (*print_calc)(position, i, c.Z, c.Rate, c.Games)
		// text_io1.LogInfo("(UCT Calculating...) %2d:z=%s,rate=%.4f,games=%3d\n", i, p.GetGtpZ(position, c.Z), c.Rate, c.Games)
	}

	// 結果
	var bestZ = pN.Children[bestI].Z
	// FIXME: (*print_calc_fin)(position, bestZ, pN.Children[bestI].Rate, max, all_playouts.AllPlayouts, node_struct.NodeNum)
	//text_io1.LogInfo("(UCT Calculated    ) bestZ=%s,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
	//	p.GetGtpZ(position, bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ, pN.Children[bestI].Rate
}

// SearchUct - 再帰関数。 GetBestZByUct() から呼び出されます
func SearchUct(
	observerGameSettingsModel *gamesettingsmodel.ObserverGameSettingsModel,
	position *position.Position,
	color color.Color,
	nodeN int) int {

	var pN = &node_struct.Nodes[nodeN]
	var c *child.Child

	for {
		var selectI = selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		var z = c.Z

		var err = position.PutStone(z, color)
		if err == 0 { // 石が置けたなら
			break
		}

		c.Z = uct_struct.IllegalZ
		// code.Console.Debug("ILLEGAL:z=%04d\n", GetZ4(z))
	}

	var winner int // 手番が勝ちなら1、引分けなら0、手番の負けなら-1 としてください
	if c.Games <= 0 {
		winner = -playout.Playout(observerGameSettingsModel, position, color.Flip(), all_playouts.GettingOfWinnerOnDuringUCTPlayout, all_playouts.IsDislike)
	} else {
		if c.Next == uct_struct.NodeEmpty {
			c.Next = node_struct.CreateNode(position)
		}
		winner = -SearchUct(observerGameSettingsModel, position, color.Flip(), c.Next)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(winner)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return winner
}

// 一番良い UCB を選びます。 SearchUct から呼び出されます。
func selectBestUcb(nodeN int) int {
	var pN = &node_struct.Nodes[nodeN]
	var selectI = -1
	var maxUcb = -999.0
	var ucb = 0.0
	for i := 0; i < pN.ChildNum; i++ {
		var c = &pN.Children[i]
		if c.Z == uct_struct.IllegalZ {
			continue
		}
		if c.Games == 0 {
			ucb = 10000.0 + 32768.0*rand.Float64()
		} else {
			ucb = c.Rate + 1.0*math.Sqrt(math.Log(float64(pN.ChildGameSum))/float64(c.Games))
		}
		if maxUcb < ucb {
			maxUcb = ucb
			selectI = i
		}
	}

	// 異常終了
	if selectI == -1 {
		coding_obj.Console.Error("Err! select\n")
		os.Exit(0)
	}

	return selectI
}
