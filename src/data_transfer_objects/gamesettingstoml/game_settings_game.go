package gamesettingstoml

import (
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/komi_float"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/moves_num"
)

// Game - [Game] テーブル
type Game struct {
	Komi      float32
	BoardSize int8
	MaxMoves  int16
	BoardData string
}

func (game *Game) GetKomi() komi_float.KomiFloat {
	return komi_float.KomiFloat(game.Komi)
}

func (game *Game) GetBoardSize() int {
	return int(game.BoardSize)
}

func (game *Game) GetMaxMoves() moves_num.MovesNum {
	return moves_num.MovesNum(game.MaxMoves)
}
