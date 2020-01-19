package palette

import h "github.com/solympe/Golang_Training/pkg/patterns/pattern-memento/history"

// represents palette struct
type palette struct {
	color   string
	history h.HistoryService
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
func (p *palette) SetHistory(historyIn h.HistoryService) {
	p.history = historyIn
}

// NewPaletteChanger returns new copy of palette
func NewPaletteChanger() PaletteChanger {
	return &palette{}
}
