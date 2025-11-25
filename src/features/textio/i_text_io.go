package textio

type ITextIO interface {
	// SendCommand - コマンド出力
	SendCommand(command string)

	// ReceivedCommand - コマンド受信
	ReceivedCommand(command string)

	// LogInfo - 情報ログ
	LogInfo(info string)
}
