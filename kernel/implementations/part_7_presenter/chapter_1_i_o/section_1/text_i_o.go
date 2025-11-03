package text_i_o

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

func (t *TextIO) GoCommand(command string) {
	//fmt.Print(command)
	code.Gtp.Print("= \n\n")
}
