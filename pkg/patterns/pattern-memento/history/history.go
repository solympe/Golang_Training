package history

type history struct {
	historyStack []string
}

// SaveColor adds new color into the stack
func (h *history) SaveColor(color string) {
	h.historyStack = append(h.historyStack, color)
}

// DeleteColor deletes from stack
func (h *history) DeleteColor() {
	h.historyStack = h.historyStack[:len(h.historyStack)-1]
}

// ExtractColor return first color from stack
func (h *history) ExtractColor() string {
	if len(h.historyStack) == 0 {
		return "history is empty"
	}
	color := h.historyStack[len(h.historyStack)-1]
	h.DeleteColor()
	return color
}

// NewHistoryService returns new copy of history
func NewHistoryService() HistoryService {
	return &history{historyStack: []string{}}
}
