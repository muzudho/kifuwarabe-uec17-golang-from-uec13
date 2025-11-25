package models

// Entities
import color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models/color"

// Ren - 連
type Ren struct {
	// LibertyArea - 呼吸点の数
	LibertyArea int
	// StoneArea - 石の数
	StoneArea int
	// Color - 石の色
	Color color.Color
}

func NewRen(libertyArea int, stoneArea int, color color.Color) *Ren {
	var ren1 = new(Ren)
	ren1.LibertyArea = libertyArea
	ren1.StoneArea = stoneArea
	ren1.Color = color
	return ren1
}
