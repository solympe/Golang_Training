package history

// HistoryService ...
type HistoryService interface {
	ExtractColor() string
	SaveColor(color string)
}
