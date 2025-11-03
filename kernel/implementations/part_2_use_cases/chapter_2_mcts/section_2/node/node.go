package node

import "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/child"

// Node - ノード。
type Node struct {
	ChildNum     int
	Children     []child.Child
	ChildGameSum int
}
