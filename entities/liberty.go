package entities

// Entities
import color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"

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
	var ren = new(Ren)
	ren.LibertyArea = libertyArea
	ren.StoneArea = stoneArea
	ren.Color = color
	return ren
}
