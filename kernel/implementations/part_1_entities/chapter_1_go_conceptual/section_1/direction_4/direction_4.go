package direction_4

// Direction4 - ４方向（東、西、南、北）の番地。水平方向、垂直方向の順で並べた
type Direction4 int

const (
	East Direction4 = iota
	North
	West
	South
)
