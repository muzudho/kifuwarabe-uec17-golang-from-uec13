package i_text_io

type ITextIO interface {
	// SendCommand - コマンド出力
	SendCommand(command string)

	// ReceivedCommand - コマンド受信
	ReceivedCommand(command string)
}
