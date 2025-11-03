package node_struct

import (
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/uct_struct"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/node"
)

// Nodes -ノードの配列？
var Nodes = [uct_struct.NodeMax]node.Node{}

// NodeNum - ノード数？
var NodeNum = 0
