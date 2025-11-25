package gamesettings

import (
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"
)

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

type ReadonlyGameSettingsModel struct {
	// boardSize - 何路盤
	boardSize int
	// komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
	komi models.KomiFloat
	// maxMovesNum - 上限手数
	maxMovesNum      models.MovesNum
	directions4Array [4]models.Point
}

func NewReadonlyGameSettingsModel(boardSize int, komi models.KomiFloat, maxMovesNum models.MovesNum) *ReadonlyGameSettingsModel {
	return &ReadonlyGameSettingsModel{
		boardSize:   boardSize,
		komi:        komi,
		maxMovesNum: maxMovesNum,
		// directions4Array - ４方向（東、西、南、北）の番地。水平方向、垂直方向の順で並べた。Directions4型の順番に対応
		directions4Array: [4]models.Point{1, -1, models.Point(boardSize + 2), models.Point(-(boardSize + 2))},
	}
}

// GetBoardSize - 壁無し盤の１辺の長さ
func (model1 *ReadonlyGameSettingsModel) GetBoardSize() int {
	return model1.boardSize
}

// GetBoardArea - 壁無し盤の面積
func (model1 *ReadonlyGameSettingsModel) GetBoardArea() int {
	return model1.boardSize * model1.boardSize
}

// GetSentinelWidth - 枠付きの盤の一辺の交点数
func (model1 *ReadonlyGameSettingsModel) GetSentinelWidth() int {
	return model1.boardSize + 2
}

// GetSentinelBoardArea - 壁付き盤の面積
func (model1 *ReadonlyGameSettingsModel) GetSentinelBoardArea() int {
	return model1.GetSentinelWidth() * model1.GetSentinelWidth()
}

func (model1 *ReadonlyGameSettingsModel) GetKomi() models.KomiFloat {
	return model1.komi
}

func (model1 *ReadonlyGameSettingsModel) GetMaxMovesNum() models.MovesNum {
	return model1.maxMovesNum
}

func (model1 *ReadonlyGameSettingsModel) GetDirections4Array() *[4]models.Point {
	return &model1.directions4Array
}
