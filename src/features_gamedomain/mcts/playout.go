package mcts

import (
	"math/rand"

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/models/color"
)

// Playout - 最後まで石を打ちます。得点を返します
// * `getWinner` - 地計算
//
// # Returns
//
// 手番が勝ったら 1、引分けなら 0、 相手が勝ったら -1
func Playout(
	readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel,
	position1 *position.Position,
	turnColor color.Color,
	getWinner *func(color.Color) int,
	isDislike *func(color.Color, models.Point) bool) int {

	AllPlayouts++

	var color = turnColor
	var previousZ models.Point = 0
	var boardMax = readonlyGameSettingsModel.GetSentinelBoardArea()

	var playoutTrialCount = PlayoutTrialCount
	for trial := 0; trial < playoutTrialCount; trial++ {
		var emptyArray = make([]models.Point, boardMax)
		var emptyLength int // 残りの空点の数
		var z models.Point

		// 空点を記憶します
		// TODO 空点を差分更新できないか？ 毎回スキャンは重くないか？
		var onPoint = func(z models.Point) {
			if position1.IsEmpty(z) { // 空点なら
				emptyArray[emptyLength] = z
				emptyLength++
			}
		}
		position1.IterateWithoutWall(onPoint)

		var r = 0
		var dislikeZ = models.Pass
		var randomPigeonX = GetRandomPigeonX(emptyLength) // 見切りを付ける試行回数を算出
		var i int
		for i = 0; i < randomPigeonX; i++ {
			if emptyLength == 0 { // 空点が無ければパスします
				z = models.Pass
			} else {

				// UEC: 改造ポイント
				// 手早く打つために、空点の一部だけをピックアップして選びます
				var shurinkenEmptyLength = emptyLength
				// var minLength = 40
				// if minLength < shurinkenEmptyLength {
				// 	shurinkenEmptyLength = int(shurinkenEmptyLength / 2)
				// }

				r = rand.Intn(shurinkenEmptyLength) // 空点を適当に選びます
				z = emptyArray[r]
			}

			var err = position1.PutStone(readonlyGameSettingsModel, z, color)
			if err == 0 { // 石が置けたか、パスなら

				if z == models.Pass || // パスか、
					!(*isDislike)(color, z) { // 石を置きたくないわけでなければ
					break // 確定
				}

				dislikeZ = z // 候補が無かったときに使います
			}

			// 石を置かなかったら、その選択肢は、最後尾の要素で置換し、最後尾の要素を消します
			emptyArray[r] = emptyArray[emptyLength-1]
			emptyLength--
		}
		if i == randomPigeonX {
			z = dislikeZ
		}

		// テストのときは棋譜を残します
		if FlagTestPlayout != 0 {
			position1.Record[position1.MovesNum].SetZ(z)
			position1.MovesNum++
		}

		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z

		color = color.Flip()
	}

	return (*getWinner)(turnColor)
}
