package child

import "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"

// Child - 子。
type Child struct {
	// table index. 盤の交点の配列のインデックス。
	Z     point.Point
	Games int     // UCT検索をした回数？
	Rate  float64 // 勝率
	Next  int     // 配列のインデックス
}
