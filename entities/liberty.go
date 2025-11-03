package entities

// Entities
import (
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	ren "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_2/ren"
)

func NewRen(libertyArea int, stoneArea int, color color.Color) *ren.Ren {
	var ren1 = new(ren.Ren)
	ren1.LibertyArea = libertyArea
	ren1.StoneArea = stoneArea
	ren1.Color = color
	return ren1
}
