package mctsimpl

import (
	"fmt"
	"time"

	// Entities

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamerecord"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamerecordpresenter"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamerecordusecase"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models/color"

	// User Cases
	mcts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mcts"
	uct "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mcts/uct"

	// Presenters

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"

	// Interfaces
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
)

// SelfPlay - コンピューター同士の対局。
func SelfPlay(text_io1 i_text_io.ITextIO, readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position1 *position.Position) {
	logger.Console.Trace("# GoGo SelfPlay 自己対局開始☆（＾～＾）\n")

	var color = color.Black

	for {
		var z = GetComputerMoveDuringSelfPlay(text_io1, readonlyGameSettingsModel, position1, color)

		var recItem = new(gamerecord.GameRecordItem)
		recItem.Z = z
		position1.PutStoneOnRecord(readonlyGameSettingsModel, z, color, recItem)

		logger.Console.Print("z=%s,color=%d", gamerecordusecase.GetGtpZ(readonlyGameSettingsModel, position1, z), color) // テスト

		// p.PrintCheckBoard(readonlyGameSettingsModel, position1)                                        // テスト
		gamerecordpresenter.PrintBoard(readonlyGameSettingsModel, position1, position1.MovesNum)

		// パスで２手目以降で棋譜の１つ前（相手）もパスなら終了します。
		if z == models.Pass && 1 < position1.MovesNum && position1.Record[position1.MovesNum-2].GetZ() == models.Pass {
			break
		}
		// 自己対局は400手で終了します。
		if 400 < position1.MovesNum {
			break
		} // too long
		color = color.Flip()
	}

	gamerecordusecase.PrintSgf(readonlyGameSettingsModel, position1, position1.MovesNum, position1.Record)
}

// GetComputerMoveDuringSelfPlay - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveDuringSelfPlay(text_io1 i_text_io.ITextIO, readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position1 *position.Position, color color.Color) models.Point {

	var start = time.Now()
	mcts.AllPlayouts = 0

	var z, winRate = uct.GetBestZByUct(
		readonlyGameSettingsModel,
		position1,
		color,
		gamerecordpresenter.CreatePrintingOfCalc(text_io1, readonlyGameSettingsModel),
		gamerecordpresenter.CreatePrintingOfCalcFin(text_io1, readonlyGameSettingsModel))

	var sec = time.Since(start).Seconds()
	text_io1.LogInfo(fmt.Sprintf("(GetComputerMoveDuringSelfPlay) %.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(mcts.AllPlayouts)/sec, position1.GetZ4(readonlyGameSettingsModel, z), winRate, position1.MovesNum, color, mcts.AllPlayouts))
	return z
}
