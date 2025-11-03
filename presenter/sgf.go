package presenter

import (
	code "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/entities"
	game_record_item "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/game_record_item"
)

// PrintSgf - SGF形式の棋譜表示。
func PrintSgf(position *e.Position, movesNum int, gameRecord []*game_record_item.GameRecordItem) {
	var boardSize = e.BoardSize

	code.Console.Print("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", boardSize, e.Komi)
	for i := 0; i < movesNum; i++ {
		var z = gameRecord[i].GetZ()
		var y = int(z) / e.SentinelWidth
		var x = int(z) - y*e.SentinelWidth
		var sStone = [2]string{"B", "W"}
		code.Console.Print(";%s", sStone[i&1])
		if z == 0 {
			code.Console.Print("[]")
		} else {
			code.Console.Print("[%c%c]", x+'a'-1, y+'a'-1)
		}
		if ((i + 1) % 10) == 0 {
			code.Console.Print("\n")
		}
	}
	code.Console.Print(")\n")
}
