package mctsimpl

import (
	"math"
	"time"

	// Entities

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamerecord"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamerecordpresenter"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mcts"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models/color"

	// Use Cases
	uct "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mcts/uct"

	// Presenters

	// Interfaces
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
)

// PlayComputerMoveLesson09a - コンピューター・プレイヤーの指し手。 SelfPlay, RunGtpEngine から呼び出されます。
func PlayComputerMoveLesson09a(
	text_io1 i_text_io.ITextIO,
	readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel,
	position1 *position.Position,
	color1 color.Color) models.Point {

	var st1 = time.Now()
	mcts.AllPlayouts = 0

	var z1, winRate1 = uct.GetBestZByUct(
		readonlyGameSettingsModel,
		position1,
		color1,
		gamerecordpresenter.CreatePrintingOfCalc(text_io1, readonlyGameSettingsModel),
		gamerecordpresenter.CreatePrintingOfCalcFin(text_io1, readonlyGameSettingsModel))

	if 1 < position1.MovesNum && // 初手ではないとして
		position1.Record[position1.MovesNum-1].GetZ() == 0 && // １つ前の手がパスで
		0.95 <= math.Abs(winRate1) { // 95%以上の確率で勝ちか負けなら
		// こちらもパスします
		return 0
	}

	var sec1 = time.Since(st1).Seconds()
	// FIXME: text_io1.LogInfo(fmt.Sprintf("%.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,movesNum=%d,color=%d,playouts=%d\n",
	//		sec1, float64(mcts.AllPlayouts)/sec1, position1.GetZ4(z1), winRate1, position1.MovesNum, color1, mcts.AllPlayouts))

	var recItem1 = new(gamerecord.GameRecordItem)
	recItem1.Z = z1
	recItem1.Time = sec1
	position1.PutStoneOnRecord(readonlyGameSettingsModel, z1, color1, recItem1)

	// FIXME: board_view.PrintBoard(readonlyGameSettingsModel, position1, position1.MovesNum)

	return z1
}
