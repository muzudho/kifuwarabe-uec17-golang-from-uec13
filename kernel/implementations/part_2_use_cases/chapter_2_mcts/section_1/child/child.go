package child

import "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"

// Child - 子。
type Child struct {
	// table index. 盤の交点の配列のインデックス。
	Z     models.Point
	Games int     // UCT検索をした回数？
	Rate  float64 // 勝率
	Next  int     // 配列のインデックス
}
