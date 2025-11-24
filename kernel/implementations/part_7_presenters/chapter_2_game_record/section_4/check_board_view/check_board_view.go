package check_board_view

import (
	"strings"

	board_view "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_2_game_record/section_3/board_view"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	logger "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
)

// " 0" - 空点
// " 1" - 黒石
var numberLabels = [2]string{" 0", " 1"}

// PrintCheckBoard - チェックボードを描画。
func PrintCheckBoard(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position1 *position.Position) {

	var b = &strings.Builder{}
	b.Grow(board_view.Sz8k)

	var boardSize = readonlyGameSettingsModel.GetBoardSize()

	// Header
	b.WriteString("\n   ")
	for x := 0; x < boardSize; x++ {
		b.WriteString(board_view.LabelOfColumns[x+1])
	}
	b.WriteString("\n  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	// Body
	for y := 0; y < boardSize; y++ {
		b.WriteString(board_view.LabelOfRows[y+1])
		b.WriteString("|")
		for x := 0; x < boardSize; x++ {
			var z = position1.GetZFromXy(readonlyGameSettingsModel, x, y)
			var number = position1.CheckAt(z)
			b.WriteString(numberLabels[number])
		}
		b.WriteString(" |\n")
	}

	// Footer
	b.WriteString("  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	logger.Console.Print(b.String())
}
