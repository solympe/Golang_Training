package historyService

// HistoryService represents history interface
type HistoryService interface {
	ExtractColor() string
	SaveColor(color string)
}
