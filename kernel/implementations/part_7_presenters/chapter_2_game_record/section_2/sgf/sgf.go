package sgf

import (
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
	game_rule_settings "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_2_rule_settings/section_1/game_rule_settings"
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_3_position/section_1"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_0_logger/section_1/coding_obj"
)

// PrintSgf - SGF形式の棋譜表示。
func PrintSgf(position *position.Position, movesNum int, gameRecord []*game_record_item.GameRecordItem) {
	var boardSize = game_rule_settings.BoardSize

	coding_obj.Console.Print("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", boardSize, game_rule_settings.Komi)
	for i := 0; i < movesNum; i++ {
		var z = gameRecord[i].GetZ()
		var y = int(z) / game_rule_settings.SentinelWidth
		var x = int(z) - y*game_rule_settings.SentinelWidth
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
