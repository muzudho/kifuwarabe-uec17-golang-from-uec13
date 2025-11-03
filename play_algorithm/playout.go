package play_algorithm

import (
	"math/rand"

	// Entities
	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	game_rule_settings "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_2_rule_settings/section_1/game_rule_settings"
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_3_position/section_1/position"
	parameter_adjustment "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/parameter_adjustment"
	all_playouts "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_2/all_playouts"
)

// Playout - 最後まで石を打ちます。得点を返します
// * `getWinner` - 地計算
//
// # Returns
//
// 手番が勝ったら 1、引分けなら 0、 相手が勝ったら -1
func Playout(
	position *position.Position,
	turnColor color.Color,
	getWinner *func(color.Color) int,
	isDislike *func(color.Color, point.Point) bool) int {

	all_playouts.AllPlayouts++

	var color = turnColor
	var previousZ point.Point = 0
	var boardMax = game_rule_settings.SentinelBoardArea

	var playoutTrialCount = parameter_adjustment.PlayoutTrialCount
	for trial := 0; trial < playoutTrialCount; trial++ {
		var empty = make([]point.Point, boardMax)
		var emptyNum int
		var z point.Point

		// TODO 空点を差分更新できないか？ 毎回スキャンは重くないか？
		// 空点を記憶します
		var onPoint = func(z point.Point) {
			if position.IsEmpty(z) { // 空点なら
				empty[emptyNum] = z
				emptyNum++
			}
		}
		position.IterateWithoutWall(onPoint)

		var r = 0
		var dislikeZ = point.Pass
		var randomPigeonX = parameter_adjustment.GetRandomPigeonX(emptyNum) // 見切りを付ける試行回数を算出
		var i int
		for i = 0; i < randomPigeonX; i++ {
			if emptyNum == 0 { // 空点が無ければパスします
				z = point.Pass
			} else {
				r = rand.Intn(emptyNum) // 空点を適当に選びます
				z = empty[r]
			}

			var err = position.PutStone(z, color)
			if err == 0 { // 石が置けたか、パスなら

				if z == point.Pass || // パスか、
					!(*isDislike)(color, z) { // 石を置きたくないわけでなければ
					break // 確定
				}

				dislikeZ = z // 候補が無かったときに使います
			}

			// 石を置かなかったら、その選択肢は、最後尾の要素で置換し、最後尾の要素を消します
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if i == randomPigeonX {
			z = dislikeZ
		}

		// テストのときは棋譜を残します
		if all_playouts.FlagTestPlayout != 0 {
			position.Record[position.MovesNum].SetZ(z)
			position.MovesNum++
		}

		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z

		color = color.Flip()
	}

	return (*getWinner)(turnColor)
}
