package logger

// StderrLogger - エラー出力と一緒に使うロガー。
type StderrLogger struct {
	tracePath  string
	debugPath  string
	infoPath   string
	noticePath string
	warnPath   string
	errorPath  string
	fatalPath  string
	printPath  string
}

func (logger1 *StderrLogger) SetPath(
	tracePath string,
	debugPath string,
	infoPath string,
	noticePath string,
	warnPath string,
	errorPath string,
	fatalPath string,
	printPath string) {

	logger1.tracePath = tracePath
	logger1.debugPath = debugPath
	logger1.infoPath = infoPath
	logger1.noticePath = noticePath
	logger1.warnPath = warnPath
	logger1.errorPath = errorPath
	logger1.fatalPath = fatalPath
	logger1.printPath = printPath
}

// Trace - ログファイルに書き込みます。
func (logger1 *StderrLogger) Trace(text string, args ...interface{}) {
	write(logger1.tracePath, text, args...)
}

// Debug - ログファイルに書き込みます。
func (logger1 *StderrLogger) Debug(text string, args ...interface{}) {
	write(logger1.debugPath, text, args...)
}

// Info - ログファイルに書き込みます。
func (logger1 *StderrLogger) Info(text string, args ...interface{}) {
	write(logger1.infoPath, text, args...)
}

// Notice - ログファイルに書き込みます。
func (logger1 *StderrLogger) Notice(text string, args ...interface{}) {
	write(logger1.noticePath, text, args...)
}

// Warn - ログファイルに書き込みます。
func (logger1 *StderrLogger) Warn(text string, args ...interface{}) {
	write(logger1.warnPath, text, args...)
}

// Error - ログファイルに書き込みます。
func (logger1 *StderrLogger) Error(text string, args ...interface{}) {
	write(logger1.errorPath, text, args...)
}

// Fatal - ログファイルに書き込みます。
func (logger1 *StderrLogger) Fatal(text string, args ...interface{}) {
	write(logger1.fatalPath, text, args...)
}

// Print - ログファイルに書き込みます。 StdoutLogWriter から呼び出してください。
func (logger1 *StderrLogger) Print(text string, args ...interface{}) {
	write(logger1.printPath, text, args...)
}
