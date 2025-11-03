package text_io

import "github.com/muzudho/kifuwarabe-uec17-golang-from-uec13/kernel/implementations/part_7_presenters/chapter_0_logger/section_1/coding_obj"

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
	coding_obj.Gtp.Print("= \n\n")
}

func (t *TextIO) ReceivedCommand(command string) {
	coding_obj.Gtp.Log(command + "\n")
	coding_obj.ConsoleLog.Notice(command + "\n")
}

func (t *TextIO) LogInfo(info string) {
	coding_obj.Console.Info(info)
}
