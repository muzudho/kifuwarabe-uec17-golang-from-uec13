package color

// Color - 石の色
type Color int

const (
	// None - 空点
	None Color = iota
	// Black - 黒石
	Black
	// White - 白石
	White
	// Wall - 壁
	Wall
)
