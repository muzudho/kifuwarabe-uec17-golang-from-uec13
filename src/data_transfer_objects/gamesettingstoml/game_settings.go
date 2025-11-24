package gamesettingstoml

// GameSettingsFile - Tomlファイル
type GameSettingsFile struct {
	// Nngs - No Name Go Server 接続設定
	Nngs Nngs
	// Game - 対局設定
	Game Game
}
