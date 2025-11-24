package z_code

import (
	"fmt"
	"strings"

	point "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/point"
	uct_struct "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_2_use_cases/chapter_2_mcts/section_1/uct_struct"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/position"
)

// GetGtpZ - XY座標をアルファベット、数字で表したもの。 例: Q10
func GetGtpZ(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position1 *position.Position, z point.Point) string {
	switch z {
	case 0:
		return "PASS"
	case uct_struct.IllegalZ:
		return "ILLEGAL" // GTP の仕様外です
	}

	var y = int(z) / readonlyGameSettingsModel.GetSentinelWidth()
	var x = int(z) % readonlyGameSettingsModel.GetSentinelWidth()

	// 筋が25（'Z'）より大きくなることは想定していません
	var alphabet_x = 'A' + x - 1
	if alphabet_x >= 'I' {
		alphabet_x++
	}

	// code.Console.Debug("y=%d x=%d z=%d alphabet_x=%d alphabet_x(char)=%c\n", y, x, z, alphabet_x, alphabet_x)

	return fmt.Sprintf("%c%d", alphabet_x, y)
}

// GetZFromGtp - GTPの座標符号を z に変換します
// * `gtp_z` - 最初の１文字はアルファベット、２文字目（あれば３文字目）は数字と想定。 例: q10
func GetZFromGtp(readonlyGameSettingsModel *gamesettings.ReadonlyGameSettingsModel, position1 *position.Position, gtp_z string) point.Point {
	gtp_z = strings.ToUpper(gtp_z)

	if gtp_z == "PASS" {
		return 0
	}

	// 筋
	var x = gtp_z[0] - 'A' + 1
	if gtp_z[0] >= 'I' {
		x--
	}

	// 段
	var y = int(gtp_z[1] - '0')
	if 2 < len(gtp_z) {
		y *= 10
		y += int(gtp_z[2] - '0')
	}

	// インデックス
	var z = position1.GetZFromXy(readonlyGameSettingsModel, int(x)-1, y-1)
	// code.Console.Trace("# x=%d y=%d z=%d z4=%04d\n", x, y, z, position1.GetZ4(z))
	return z
}
