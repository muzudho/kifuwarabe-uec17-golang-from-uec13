package gamesettingsmodel

import (
	komi_float "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/komi_float"
	moves_num "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/moves_num"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
)

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

type ReadonlyGameSettingsModel struct {
	// boardSize - 何路盤
	boardSize int
	// komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
	komi komi_float.KomiFloat
	// maxMovesNum - 上限手数
	maxMovesNum      moves_num.MovesNum
	directions4Array [4]point.Point
}

func NewReadonlyGameSettingsModel(boardSize int, komi komi_float.KomiFloat, maxMovesNum moves_num.MovesNum) *ReadonlyGameSettingsModel {
	return &ReadonlyGameSettingsModel{
		boardSize:   boardSize,
		komi:        komi,
		maxMovesNum: maxMovesNum,
		// directions4Array - ４方向（東、西、南、北）の番地。水平方向、垂直方向の順で並べた。Directions4型の順番に対応
		directions4Array: [4]point.Point{1, -1, point.Point(boardSize + 2), point.Point(-(boardSize + 2))},
	}
}

// GetBoardSize - 壁無し盤の１辺の長さ
func (model *ReadonlyGameSettingsModel) GetBoardSize() int {
	return model.boardSize
}

// GetBoardArea - 壁無し盤の面積
func (model *ReadonlyGameSettingsModel) GetBoardArea() int {
	return model.boardSize * model.boardSize
}

// GetSentinelWidth - 枠付きの盤の一辺の交点数
func (model *ReadonlyGameSettingsModel) GetSentinelWidth() int {
	return model.boardSize + 2
}

// GetSentinelBoardArea - 壁付き盤の面積
func (model *ReadonlyGameSettingsModel) GetSentinelBoardArea() int {
	return model.GetSentinelWidth() * model.GetSentinelWidth()
}

func (model *ReadonlyGameSettingsModel) GetKomi() komi_float.KomiFloat {
	return model.komi
}

func (model *ReadonlyGameSettingsModel) GetMaxMovesNum() moves_num.MovesNum {
	return model.maxMovesNum
}

func (model *ReadonlyGameSettingsModel) GetDirections4Array() *[4]point.Point {
	return &model.directions4Array
}
