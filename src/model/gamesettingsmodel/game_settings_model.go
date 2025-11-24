package gamesettingsmodel

import (
	// Entities
	komi_float "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/komi_float"
	moves_num "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/moves_num"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
)

// // GameSettingsModel - 対局設定
// type GameSettingsModel struct {
// }

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

// Komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
var Komi komi_float.KomiFloat

// MaxMovesNum - 上限手数
var MaxMovesNum moves_num.MovesNum

// Directions4Array - ４方向（東、西、南、北）の番地。水平方向、垂直方向の順で並べた。Directions4型の順番に対応
var Directions4Array = [4]point.Point{1, -1, 9, -9}
