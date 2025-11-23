package node_struct

import (
	"os"

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_3_position"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/child"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/uct_struct"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/node"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_0_logger/section_1/coding_obj"
)

// Nodes -ノードの配列？
var Nodes = [uct_struct.NodeMax]node.Node{}

// NodeNum - ノード数？
var NodeNum = 0

// CreateNode - ノード作成。 searchUctV8, GetBestZByUct, searchUctLesson09 から呼び出されます。
func CreateNode(position *position.Position) int {

	if NodeNum == uct_struct.NodeMax {
		coding_obj.Console.Error("node over Err\n")
		os.Exit(0)
	}
	var pN = &Nodes[NodeNum]
	pN.ChildNum = 0
	pN.Children = make([]child.Child, position.UctChildrenSize())
	pN.ChildGameSum = 0

	var onPoint = func(z point.Point) {
		if position.IsEmpty(z) { // 空点なら
			addChild(pN, z)
		}
	}
	position.IterateWithoutWall(onPoint)

	addChild(pN, 0)
	NodeNum++
	return NodeNum - 1
}

// CreateNode から呼び出されます。
func addChild(pN *node.Node, z point.Point) {
	var n = pN.ChildNum
	pN.Children[n].Z = z
	pN.Children[n].Games = 0
	pN.Children[n].Rate = 0.0
	pN.Children[n].Next = uct_struct.NodeEmpty
	pN.ChildNum++
}
