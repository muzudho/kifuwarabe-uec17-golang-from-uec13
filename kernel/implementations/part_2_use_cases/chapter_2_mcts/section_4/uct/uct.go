package uct

import (
	"math"
	"math/rand"
	"os"

	// Entities

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mcts"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mcts/nodestruct"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mcts/uctstruct"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models/color"
)

// GetBestZByUct - Lesson08,09,09aで使用。 一番良いUCTである着手を選びます。 GetComputerMoveDuringSelfPlay などから呼び出されます。
//
// # Return
// (bestZ int, winRate float64)
func GetBestZByUct(
	readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel,
	position1 *position.Position,
	color color.Color,
	print_calc *func(*position.Position, int, models.Point, float64, int),
	print_calc_fin *func(*position.Position, models.Point, float64, int, int, int)) (models.Point, float64) {

	// UCT計算フェーズ
	nodestruct.NodeNum = 0 // カウンターリセット
	var next = nodestruct.CreateNode(position1)
	var uctLoopCount = mcts.UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		// 一時記憶
		var copiedPosition = position1.CopyPosition(readonlyGameSettingsModel)

		SearchUct(readonlyGameSettingsModel, position1, color, next)

		// 復元
		position1.ImportPosition(copiedPosition)
	}

	// ベスト値検索フェーズ
	var bestI = -1
	var max = -999
	var pN = &nodestruct.Nodes[next]
	for i := 0; i < pN.ChildNum; i++ {
		var c = &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		// FIXME: (*print_calc)(position1, i, c.Z, c.Rate, c.Games)
		// text_io1.LogInfo("(UCT Calculating...) %2d:z=%s,rate=%.4f,games=%3d\n", i, p.GetGtpZ(position1, c.Z), c.Rate, c.Games)
	}

	// 結果
	var bestZ = pN.Children[bestI].Z
	// FIXME: (*print_calc_fin)(position1, bestZ, pN.Children[bestI].Rate, max, mcts.AllPlayouts, nodestruct.NodeNum)
	//text_io1.LogInfo("(UCT Calculated    ) bestZ=%s,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
	//	p.GetGtpZ(position1, bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ, pN.Children[bestI].Rate
}

// SearchUct - 再帰関数。 GetBestZByUct() から呼び出されます
func SearchUct(
	readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel,
	position1 *position.Position,
	color color.Color,
	nodeN int) int {

	var pN = &nodestruct.Nodes[nodeN]
	var c *mcts.Child

	for {
		var selectI = selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		var z = c.Z

		var err = position1.PutStone(readonlyGameSettingsModel, z, color)
		if err == 0 { // 石が置けたなら
			break
		}

		c.Z = uctstruct.IllegalZ
		// code.Console.Debug("ILLEGAL:z=%04d\n", GetZ4(z))
	}

	var winner int // 手番が勝ちなら1、引分けなら0、手番の負けなら-1 としてください
	if c.Games <= 0 {
		winner = -mcts.Playout(readonlyGameSettingsModel, position1, color.Flip(), mcts.GettingOfWinnerOnDuringUCTPlayout, mcts.IsDislike)
	} else {
		if c.Next == uctstruct.NodeEmpty {
			c.Next = nodestruct.CreateNode(position1)
		}
		winner = -SearchUct(readonlyGameSettingsModel, position1, color.Flip(), c.Next)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(winner)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return winner
}

// 一番良い UCB を選びます。 SearchUct から呼び出されます。
func selectBestUcb(nodeN int) int {
	var pN = &nodestruct.Nodes[nodeN]
	var selectI = -1
	var maxUcb = -999.0
	var ucb = 0.0
	for i := 0; i < pN.ChildNum; i++ {
		var c = &pN.Children[i]
		if c.Z == uctstruct.IllegalZ {
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
		logger.Console.Error("Err! select\n")
		os.Exit(0)
	}

	return selectI
}
