package entities

import (
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/komi_float"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/moves_num"
)

// Entities

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

func SetBoardSize(boardSize int) {
	BoardSize = boardSize
	BoardArea = BoardSize * BoardSize
	SentinelWidth = BoardSize + 2
	SentinelBoardArea = SentinelWidth * SentinelWidth
}

// BoardSize - 何路盤
var BoardSize int

// BoardArea - 壁無し盤の面積
var BoardArea int

// SentinelWidth - 枠付きの盤の一辺の交点数
var SentinelWidth int

// SentinelBoardArea - 壁付き盤の面積
var SentinelBoardArea int

// Komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
var Komi komi_float.KomiFloat

// MaxMovesNum - 上限手数
var MaxMovesNum moves_num.MovesNum

// Point - 交点の座標。壁を含む盤の左上を 0 とします
type Point int

// Pass - パス
const Pass Point = 0

// Dir4 - ４方向（東、北、西、南）の番地。初期値は仮の値。 2015年講習会サンプル、GoGo とは順序が違います
var Dir4 = [4]Point{1, -9, -1, 9}

type Direction4 int

// Dir4に対応
const (
	East Direction4 = iota
	North
	West
	South
)
