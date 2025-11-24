// 標準出力とロガーを一緒にしただけのもの
package logger

import (
	"fmt"
)

// StdoutLogWriter - 標準出力とロガーを一緒にしただけです
type StdoutLogWriter struct {
	logger1 *StdoutLogger
}

// NewStdoutLogWriter - オブジェクト作成
func NewStdoutLogWriter(logger1 *StdoutLogger) *StdoutLogWriter {
	writer := new(StdoutLogWriter)
	writer.logger1 = logger1
	return writer
}

// Print - 必ず出力します。
func (writer *StdoutLogWriter) Print(text string, args ...interface{}) {
	fmt.Printf(text, args...) // 標準出力
	// FIXME: CgfGoBan では StdErr 使ったら不具合起こす。
	//writer.logger1.Print(text, args...) // ログ
}

// Log - ログだけ
func (writer *StdoutLogWriter) Log(text string, args ...interface{}) {
	// FIXME: CgfGoBan では StdErr 使ったら不具合起こす。
	//writer.logger1.Print(text, args...) // ログ
}
