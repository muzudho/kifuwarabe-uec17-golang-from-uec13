package gamesettings

import (
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/moves_num"
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

func (game *Game) GetMaxMoves() moves_num.MovesNum {
	return moves_num.MovesNum(game.MaxMoves)
}
