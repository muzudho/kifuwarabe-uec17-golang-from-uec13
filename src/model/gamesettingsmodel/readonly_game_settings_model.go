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

// Directions4Array - ４方向（東、西、南、北）の番地。水平方向、垂直方向の順で並べた。Directions4型の順番に対応
var Directions4Array = [4]point.Point{1, -1, 9, -9}

type ObserverGameSettingsModel struct {
	// boardSize - 何路盤
	boardSize int
	// komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
	komi komi_float.KomiFloat
	// maxMovesNum - 上限手数
	maxMovesNum moves_num.MovesNum
}

func NewReadonlyGameSettingsModel(boardSize int, komi komi_float.KomiFloat, maxMovesNum moves_num.MovesNum) *ObserverGameSettingsModel {
	return &ObserverGameSettingsModel{
		boardSize:   boardSize,
		komi:        komi,
		maxMovesNum: maxMovesNum,
	}
}

// GetBoardSize - 壁無し盤の１辺の長さ
func (model *ObserverGameSettingsModel) GetBoardSize() int {
	return model.boardSize
}

// GetBoardArea - 壁無し盤の面積
func (model *ObserverGameSettingsModel) GetBoardArea() int {
	return model.boardSize * model.boardSize
}

// GetSentinelWidth - 枠付きの盤の一辺の交点数
func (model *ObserverGameSettingsModel) GetSentinelWidth() int {
	return model.boardSize + 2
}

// GetSentinelBoardArea - 壁付き盤の面積
func (model *ObserverGameSettingsModel) GetSentinelBoardArea() int {
	return model.GetSentinelWidth() * model.GetSentinelWidth()
}

func (model *ObserverGameSettingsModel) GetKomi() komi_float.KomiFloat {
	return model.komi
}

func (model *ObserverGameSettingsModel) GetMaxMovesNum() moves_num.MovesNum {
	return model.maxMovesNum
}
