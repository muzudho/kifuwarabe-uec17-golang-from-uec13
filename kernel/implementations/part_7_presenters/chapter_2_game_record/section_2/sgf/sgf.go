package sgf

import (
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	coding_obj "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"
)

// PrintSgf - SGF形式の棋譜表示。
func PrintSgf(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position *position.Position, movesNum int, gameRecord []*game_record_item.GameRecordItem) {
	var boardSize = readonlyGameSettingsModel.GetBoardSize()

	coding_obj.Console.Print("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", boardSize, readonlyGameSettingsModel.GetKomi())
	for i := 0; i < movesNum; i++ {
		var z = gameRecord[i].GetZ()
		var y = int(z) / readonlyGameSettingsModel.GetSentinelWidth()
		var x = int(z) - y*readonlyGameSettingsModel.GetSentinelWidth()
		var sStone = [2]string{"B", "W"}
		coding_obj.Console.Print(";%s", sStone[i&1])
		if z == 0 {
			coding_obj.Console.Print("[]")
		} else {
			coding_obj.Console.Print("[%c%c]", x+'a'-1, y+'a'-1)
		}
		if ((i + 1) % 10) == 0 {
			coding_obj.Console.Print("\n")
		}
	}
	coding_obj.Console.Print(")\n")
}
