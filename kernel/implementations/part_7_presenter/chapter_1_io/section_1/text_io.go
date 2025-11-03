package text_io

import code "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/coding_obj"

// TextIO - テキスト入出力
type TextIO struct {
	// ロガー
	//log1 *logger.Logger
}

//log1 *logger.Logger
func NewTextIO() *TextIO {
	var t = new(TextIO)
	//t.log1 = log1
	return t
}

func (t *TextIO) SendCommand(command string) {
	//fmt.Print(command)
	code.Gtp.Print("= \n\n")
}

func (t *TextIO) ReceivedCommand(command string) {
	code.Gtp.Log(command + "\n")
	code.ConsoleLog.Notice(command + "\n")
}
