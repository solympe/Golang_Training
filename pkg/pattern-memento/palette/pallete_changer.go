package palette

import "github.com/solympe/Golang_Training/pkg/pattern-memento/history"

// PaletteChanger represents palette interface
type PaletteChanger interface {
	SetColor(color string)
	GetColor() (color string)
	SetHistory(historyIn history.HistoryService)
	SetPreviousColor()
}
