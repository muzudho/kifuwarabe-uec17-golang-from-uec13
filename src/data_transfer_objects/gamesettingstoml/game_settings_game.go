package gamesettingstoml

// Game - [Game] テーブル
type Game struct {
	Komi      float32
	BoardSize int8
	MaxMoves  int16
	BoardData string
}
