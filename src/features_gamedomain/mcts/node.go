package mcts

// Node - ノード。
type Node struct {
	ChildNum     int
	Children     []Child
	ChildGameSum int
}
