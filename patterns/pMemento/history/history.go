package history

import (
	h "github.com/solympe/Golang_Training/patterns/pMemento/historyService"
	p "github.com/solympe/Golang_Training/patterns/pMemento/palleteChanger"
)

// history represents history struct
type history struct {
	historyStack []string
	palette      p.PaletteChanger
}

// Save color adds new color into the stack
func (h *history) SaveColor(color string) {
	h.historyStack = append(h.historyStack, color)
}

// DeleteColor deletes from stack
func (h *history) DeleteColor() {
	h.historyStack = h.historyStack[1:]
}

// ExtractColor return first color from stack
func (h *history) ExtractColor() string {
	if len(h.historyStack) == 0 {
		return ""
	}
	color := h.historyStack[0]
	h.DeleteColor()
	return color
}

// NewHistory returns new copy of history
func NewHistory() h.HistoryService {
	return &history{historyStack: []string{}}
}
