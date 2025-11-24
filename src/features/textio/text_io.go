package textio

import (
	"fmt"

	"github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/src/features/logger"
)

// TextIO - テキスト入出力
type TextIO struct {
	// ロガー
}

func NewTextIO() *TextIO {
	var t = new(TextIO)
	//t.log1 = log1
	return t
}

func (t *TextIO) SendCommand(command string) {
	fmt.Print(command)
	//logger.Gtp.Print(command)
}

func (t *TextIO) ReceivedCommand(command string) {
	logger.Gtp.Log("%s\n", command)
	logger.ConsoleLog.Notice("%s\n", command)
}

func (t *TextIO) LogInfo(info string) {
	logger.Console.Info("%s", info)
}
