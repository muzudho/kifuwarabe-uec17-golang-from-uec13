package presenter

import (
	"strconv"
	"strings"

	code "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/coding_obj"

	// Entity
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	game_rule_settings "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_2_rule_settings/section_1/game_rule_settings"
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_3_position/section_1/position"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenter/chapter_2_game_record/section_1/z_code"
)

var sz8k = 8 * 1024

// 案
//     A B C D E F G H J K L M N O P Q R S T
//   +---------------------------------------+
//  1| . . . . . . . . . . . . . . . . . . . |
//  2| . . . . . . . . . . . . . . . . . . . |
//  3| . . . . . . . . . . . . . . . . x . . |
//  4| . . . . . . . . . . . . . . . . . . . |
//  5| . . . . . . . . . . . . . . . . . . . |
//  6| . . . . . . . . . . . . . . . . . . . |
//  7| . . . . . . . . . . . . . . . . . . . |
//  8| . . . . . . . . . . . . . . . . . . . |
//  9| . . . . . . . . . . . . . . . . . . . |
// 10| . . . . . . . . . . . . . . . . . . . |
// 11| . . . . . . . . . . . . . . . . . . . |
// 12| . . . . . . . . . . . . . . . . . . . |
// 13| . . . . . . . . . . . . . . . . . . . |
// 14| . . . . . . . . . . . . . . . . . . . |
// 15| . . . . . . . . . . . . . . . . . . . |
// 16| . . . . . . . . . . . . . . . . . . . |
// 17| . . o . . . . . . . . . . . . . . . . |
// 18| . . . . . . . . . . . . . . . . . . . |
// 19| . . . . . . . . . . . . . . . . . . . |
//   +---------------------------------------+
//  KoZ=0,movesNum=999
//
// ASCII文字を使います（全角、半角の狂いがないため）
// 黒石は x 、 白石は o （ダークモードでもライトモードでも識別できるため）

// labelOfColumns - 各列の表示符号。
// 国際囲碁連盟のフォーマット
var labelOfColumns = [20]string{"xx", " A", " B", " C", " D", " E", " F", " G", " H", " J",
	" K", " L", " M", " N", " O", " P", " Q", " R", " S", " T"}

// labelOfRows - 各行の表示符号。
var labelOfRows = [20]string{" 0", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9",
	"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}

// " ." - 空点
// " x" - 黒石
// " o" - 白石
// " #" - 壁（バグ目視確認用）
var stoneLabels = [4]string{" .", " x", " o", " #"}

// " ." - 空点（バグ目視確認用）
// " x" - 黒石（バグ目視確認用）
// " o" - 白石（バグ目視確認用）
// "+-" - 壁
var leftCornerLabels = [4]string{".", "x", "o", "+"}
var horizontalEdgeLabels = [4]string{" .", " x", " o", "--"}
var rightCornerLabels = [4]string{" .", " x", " o", "-+"}
var leftVerticalEdgeLabels = [4]string{".", "x", "o", "|"}
var rightVerticalEdgeLabels = [4]string{" .", " x", " o", " |"}

// PrintBoard - 盤を描画。
func PrintBoard(position *position.Position, movesNum int) {

	var b = &strings.Builder{}
	b.Grow(sz8k)

	var boardSize = game_rule_settings.BoardSize

	// Header (numbers)
	b.WriteString("\n   ")
	for x := 0; x < boardSize; x++ {
		b.WriteString(labelOfColumns[x+1])
	}
	// Header (line)
	b.WriteString("\n  ")                                // number space
	b.WriteString(leftCornerLabels[position.ColorAt(0)]) // +
	for x := 0; x < boardSize; x++ {
		b.WriteString(horizontalEdgeLabels[position.ColorAt(point.Point(x+1))]) // --
	}
	b.WriteString(rightCornerLabels[position.ColorAt(point.Point(game_rule_settings.SentinelWidth-1))]) // -+
	b.WriteString("\n")

	// Body
	for y := 0; y < boardSize; y++ {
		b.WriteString(labelOfRows[y+1])                                                                              // number
		b.WriteString(leftVerticalEdgeLabels[position.ColorAt(point.Point((y+1)*game_rule_settings.SentinelWidth))]) // |
		for x := 0; x < boardSize; x++ {
			b.WriteString(stoneLabels[position.ColorAtXy(x, y)])
		}
		b.WriteString(rightVerticalEdgeLabels[position.ColorAt(point.Point((y+2)*game_rule_settings.SentinelWidth-1))]) // " |"
		b.WriteString("\n")
	}

	// Footer
	b.WriteString("  ") // number space
	var a = game_rule_settings.SentinelWidth * (game_rule_settings.SentinelWidth - 1)
	b.WriteString(leftCornerLabels[position.ColorAt(point.Point(a))]) // +
	for x := 0; x < boardSize; x++ {
		b.WriteString(horizontalEdgeLabels[position.ColorAt(point.Point(a+x+1))]) // --
	}
	b.WriteString(rightCornerLabels[position.ColorAt(point.Point(game_rule_settings.SentinelBoardArea-1))]) // -+
	b.WriteString("\n")

	// Info
	b.WriteString("  KoZ=")
	if position.KoZ == point.Pass {
		b.WriteString("_")
	} else {
		b.WriteString(z_code.GetGtpZ(position, position.KoZ))
	}
	if movesNum != -1 {
		b.WriteString(",movesNum=")
		b.WriteString(strconv.Itoa(movesNum))
	}
	b.WriteString("\n")

	code.Console.Print(b.String())
}
