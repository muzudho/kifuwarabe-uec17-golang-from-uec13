package gamesettings

import (
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"
)

// Game - [Game] テーブル
type Game struct {
	Komi      float32
	BoardSize int8
	MaxMoves  int16
	BoardData string
}

func (game *Game) GetKomi() models.KomiFloat {
	return models.KomiFloat(game.Komi)
}

func (game *Game) GetBoardSize() int {
	return int(game.BoardSize)
}

func (game *Game) GetMaxMoves() models.MovesNum {
	return models.MovesNum(game.MaxMoves)
}
