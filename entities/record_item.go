package entities

import "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"

// RecordItem - 棋譜の1手分
type RecordItem struct {
	// Z - 着手
	Z point.Point
	// Time - 消費時間
	Time float64
}

// SetZ - 着手
func (recItem *RecordItem) SetZ(z point.Point) {
	recItem.Z = z
}
func (recItem *RecordItem) GetZ() point.Point {
	return recItem.Z
}

// SetTime - 消費時間
func (recItem *RecordItem) SetTime(time float64) {
	recItem.Time = time
}
func (recItem *RecordItem) GetTime() float64 {
	return recItem.Time
}
