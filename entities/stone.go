package entities

// Stone - 石の色
type Stone int

const (
	// Empty - 空点
	Empty Stone = iota
	// Black - 黒石
	Black
	// White - 白石
	White
	// Wall - 壁
	Wall
)
