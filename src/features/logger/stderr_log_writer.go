// 標準出力とロガーを一緒にしただけのもの
package logger

import (
	"fmt"
	"net"
	"os"
)

// StderrLogWriter - エラー出力とロガーを一緒にしただけです
type StderrLogWriter struct {
	logger1 *StderrLogger
}

// NewStderrLogWriter - オブジェクト作成
func NewStderrLogWriter(logger1 *StderrLogger) *StderrLogWriter {
	writer := new(StderrLogWriter)
	writer.logger1 = logger1
	return writer
}

// Trace - 本番運用時にはソースコードにも残っていないような内容を書くのに使います。
func (writer *StderrLogWriter) Trace(text string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, text, args...) // エラー出力
	writer.logger1.Trace(text, args...)   // ログ
}

// Debug - 本番運用時にもデバッグを取りたいような内容を書くのに使います。
func (writer *StderrLogWriter) Debug(text string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, text, args...) // エラー出力
	writer.logger1.Debug(text, args...)   // ログ
}

// Info - 多めの情報を書くのに使います。
func (writer *StderrLogWriter) Info(text string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, text, args...) // エラー出力
	writer.logger1.Info(text, args...)    // ログ
}

// Notice - 定期的に動作確認を取りたいような、節目、節目の重要なポイントの情報を書くのに使います。
func (writer *StderrLogWriter) Notice(text string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, text, args...) // エラー出力
	writer.logger1.Notice(text, args...)  // ログ
}

// Warn - ハードディスクの残り容量が少ないなど、当面は無視できるが対応はしたいような情報を書くのに使います。
func (writer *StderrLogWriter) Warn(text string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, text, args...) // エラー出力
	writer.logger1.Warn(text, args...)    // ログ
}

// Error - 動作不良の内容や、理由を書くのに使います。
func (writer *StderrLogWriter) Error(text string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, text, args...) // エラー出力
	writer.logger1.Error(text, args...)   // ログ
}

// Fatal - 強制終了したことを伝えます。
func (writer *StderrLogWriter) Fatal(text string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, text, args...) // エラー出力
	writer.logger1.Fatal(text, args...)   // ログ
}

// Print - 必ず出力します。
func (writer *StderrLogWriter) Print(text string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, text, args...) // エラー出力
	writer.logger1.Print(text, args...)   // ログ
}

// Send - メッセージを送信します。
func (writer *StderrLogWriter) Send(conn net.Conn, text string, args ...interface{}) {
	_, err := fmt.Fprintf(conn, text, args...) // 出力先指定
	if err != nil {
		panic(err)
	}

	fmt.Printf(text, args...)           // 標準出力
	writer.logger1.Print(text, args...) // ログ
}
