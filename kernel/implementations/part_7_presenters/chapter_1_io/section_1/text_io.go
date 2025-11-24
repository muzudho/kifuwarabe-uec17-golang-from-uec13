package text_io

import (
	"fmt"

	coding_obj "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"
)

// TextIO - テキスト入出力
type TextIO struct {
	// ロガー
	//log1 *logger.Logger
}

// log1 *logger.Logger
func NewTextIO() *TextIO {
	var t = new(TextIO)
	//t.log1 = log1
	return t
}

func (t *TextIO) SendCommand(command string) {
	fmt.Print(command)
	//coding_obj.Gtp.Print(command)
}

func (t *TextIO) ReceivedCommand(command string) {
	coding_obj.Gtp.Log("%s\n", command)
	coding_obj.ConsoleLog.Notice("%s\n", command)
}

func (t *TextIO) LogInfo(info string) {
	coding_obj.Console.Info("%s", info)
}
