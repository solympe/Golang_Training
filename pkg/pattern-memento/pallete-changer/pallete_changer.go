package pallete_changer

import (
	"github.com/solympe/Golang_Training/pkg/pattern-memento/history-service"
)

// PaletteChanger represents palette interface
type PaletteChanger interface {
	SetColor(color string)
	GetColor() (color string)
	SetHistory(historyIn history_service.HistoryService)
	SetPreviousColor()
}
