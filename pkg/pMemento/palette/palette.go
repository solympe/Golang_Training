package palette

import (
	"github.com/solympe/Golang_Training/pkg/pMemento/historyService"
	"github.com/solympe/Golang_Training/pkg/pMemento/palleteChanger"
)

// represents palette struct
type palette struct {
	color   string
	history historyService.HistoryService
}

// SetColor set new color into palette
func (p *palette) SetColor(color string) {
	p.color = color
	p.history.SaveColor(color)
}

// GetColor returns actual palette color
func (p *palette) GetColor() (color string) {
	return p.color
}

// SetPreviousColor returns previous color from palette
func (p *palette) SetPreviousColor() {
	p.color = p.history.ExtractColor()

}

// SetHistory sets palettes history
func (p *palette) SetHistory(historyIn historyService.HistoryService) {
	p.history = historyIn
}

// NewPalette returns new copy of palette
func NewPalette() palleteChanger.PaletteChanger {
	return &palette{}
}
