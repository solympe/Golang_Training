package palette

import "github.com/solympe/Golang_Training/pkg/patterns/pattern-memento/history"

// PaletteChanger represents palette interface
type PaletteChanger interface {
	SetColor(color string)
	GetColor() (color string)
	SetHistory(historyIn history.HistoryService)
	SetPreviousColor()
}
