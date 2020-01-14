package palette

import (
	"github.com/solympe/Golang_Training/pkg/pattern-memento/history-service"
	"github.com/solympe/Golang_Training/pkg/pattern-memento/pallete-changer"
)

// represents palette struct
type palette struct {
	color   string
	history history_service.HistoryService
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
func (p *palette) SetHistory(historyIn history_service.HistoryService) {
	p.history = historyIn
}

// NewPalette returns new copy of palette
func NewPalette() pallete_changer.PaletteChanger {
	return &palette{}
}
