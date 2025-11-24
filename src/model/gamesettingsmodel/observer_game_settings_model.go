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

func (model *ObserverGameSettingsModel) GetBoardSize() int {
	return model.boardSize
}
