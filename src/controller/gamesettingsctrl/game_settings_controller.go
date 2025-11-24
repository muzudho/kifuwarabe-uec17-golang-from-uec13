package gamesettingsctrl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pelletier/go-toml"

	"os"

	color "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_1_entities/chapter_1_go_conceptual/section_1/color"
	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/gamesettings"
)

// LoadGameSettings - ゲーム設定ファイルを読み込みます
func LoadGameSettings(
	path string,
	onFatal func(string)) gamesettings.GameSettingsFile {

	// ファイル読込
	var fileData, err = os.ReadFile(path)
	if err != nil {
		onFatal(fmt.Sprintf("path=%s", path))
		panic(err)
	}

	// Toml解析
	var binary = []byte(string(fileData))
	var gamesettings1 = gamesettings.GameSettingsFile{}
	toml.Unmarshal(binary, &gamesettings1)

	return gamesettings1
}

// GetBoardArray - 盤上の石の色の配列
// 0: 空点
// 1: 黒石
// 2: 白石
// 3: 壁
func GetBoardArray(gamesettings1 *gamesettings.GameSettingsFile) []color.Color {
	// 最後のカンマを削除しないと、要素数が 1 多くなってしまいます
	var s = strings.TrimRight(gamesettings1.Game.BoardData, ",")
	var nodes = strings.Split(s, ",")
	var array = make([]color.Color, len(nodes))
	for i, s := range nodes {
		var s = strings.TrimSpace(s) // 前後の半角空白、改行、タブを除去
		var color1, _ = strconv.Atoi(s)
		// fmt.Printf("[%d \"%s\" %d]", i, s, color1) // デバッグ出力
		array[i] = color.Color(color1)
	}

	return array
}
