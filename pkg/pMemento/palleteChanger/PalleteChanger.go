package palleteChanger

import (
	"github.com/solympe/Golang_Training/pkg/pMemento/historyService"
)

// PaletteChanger represents palette interface
type PaletteChanger interface {
	SetColor(color string)
	GetColor() (color string)
	SetHistory(historyIn historyService.HistoryService)
	SetPreviousColor()
}
