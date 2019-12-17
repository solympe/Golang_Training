package palleteChanger

import h "github.com/solympe/Golang_Training/patterns/pMemento/historyService"

// PaletteChanger represents palette interface
type PaletteChanger interface {
	SetColor(color string)
	GetColor() (color string)
	SetHistory(historyIn h.HistoryService)
	SetPreviousColor()
}
