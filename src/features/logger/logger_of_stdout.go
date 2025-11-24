package logger

// StdoutLogger - 標準出力と一緒に使うロガー。
type StdoutLogger struct {
	printPath string
}

func (logger1 *StdoutLogger) SetPath(printPath string) {
	logger1.printPath = printPath
}

// Print - ログファイルに書き込みます。 StdoutLogWriter から呼び出してください。
func (logger1 *StdoutLogger) Print(text string, args ...interface{}) {
	write(logger1.printPath, text, args...)
}
