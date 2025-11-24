package parameter_adjustment

import (
	// Entities
	position "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities"
	gamesettingsmodel "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/model/gamesettingsmodel"
)

// プレイアウトする回数（あとで設定されます）
var PlayoutTrialCount = 0

// UCTをループする回数（あとで設定されます）
var UctLoopCount = 0

// ランダム鳩の巣仮説定数 a。およそ 18
// 面積 * 2 pi e 、つまり およそ 17 で、５００回に１回見落としがある程度、
// 面積 * (2 pi e + 1) 、 つまり およそ 18 で、１万回に１回見落としがある程度の精度（自分調べ）
var randomPigeonA = 17 // 2 * math.Pi * math.E

// ランダム鳩の巣仮説 試行回数 x
// 📖 [random-pigeon-nest-hypothesis](https://github.com/muzudho/random-pigeon-nest-hypothesis)
func GetRandomPigeonX(N int) int {
	return N * randomPigeonA
	// return int(math.Ceil(float64(N) * randomPigeonA))
}

func AdjustParameters(position *position.Position) {
	var observerGameSettingsModel = gamesettingsmodel.NewObserverGameSettingsModel(gamesettingsmodel.BoardSize)
	var boardSize = observerGameSettingsModel.GetBoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		PlayoutTrialCount = boardSize * boardSize
	}

	// UEC: 改造ポイント
	// 盤面全体を１回は選ぶことを、完璧ではありませんが、ある程度の精度でカバーします
	// UctLoopCount = GetRandomPigeonX(gamesettingsmodel.BoardArea)
	// ↓
	// 持ち時間３０分（１８００秒）。上限手数４００。１人２００。つまり、１手あたり０.９秒。
	// boardSize * 3 なら６秒。 boardSize * 5 なら１１秒。 boardSize * 4 ならピッタリ９秒。 boardSize * 3.5 なら７秒。 boardSize * 3.75 なら８秒。
	UctLoopCount = 3 // 動作テスト用［ペンキ塗り］
	//UctLoopCount = int(float64(gamesettingsmodel.BoardArea) * 3) // 試行時間６秒
	//UctLoopCount = int(float64(gamesettingsmodel.BoardArea) * 3.75) // 試行時間８秒
	// FIXME: ランダム・ピジョン（17ぐらい）を使いたいが、処理速度が遅いので、代わりに小さな数字を入れる。
	// ↓
	// 時間いっぱい考えさせてもペンキ塗りを始めるので、少なくする。
	//UctLoopCount = int(float64(gamesettingsmodel.BoardArea)*0.5) + 1
}
