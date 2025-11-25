package gamerecord

import "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"

// GameRecordItem - 棋譜の1手分
type GameRecordItem struct {
	// Z - 着手
	Z models.Point
	// Time - 消費時間
	Time float64
}

// SetZ - 着手
func (recItem *GameRecordItem) SetZ(z models.Point) {
	recItem.Z = z
}
func (recItem *GameRecordItem) GetZ() models.Point {
	return recItem.Z
}

// SetTime - 消費時間
func (recItem *GameRecordItem) SetTime(time float64) {
	recItem.Time = time
}
func (recItem *GameRecordItem) GetTime() float64 {
	return recItem.Time
}
