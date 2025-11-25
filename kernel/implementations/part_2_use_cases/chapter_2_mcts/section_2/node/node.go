package node

import mcts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features_gamedomain/mcts"

// Node - ノード。
type Node struct {
	ChildNum     int
	Children     []mcts.Child
	ChildGameSum int
}
