package gamesettingsmodel

type ObserverGameSettingsModel struct {
	// BoardSize - 何路盤
	boardSize int
}

func NewObserverGameSettingsModel(boardSize int) *ObserverGameSettingsModel {
	return &ObserverGameSettingsModel{
		boardSize: boardSize,
	}
}

// GetBoardSize - 壁無し盤の１辺の長さ
func (model *ObserverGameSettingsModel) GetBoardSize() int {
	return model.boardSize
}

// GetBoardArea - 壁無し盤の面積
func (model *ObserverGameSettingsModel) GetBoardArea() int {
	return model.boardSize * model.boardSize
}
