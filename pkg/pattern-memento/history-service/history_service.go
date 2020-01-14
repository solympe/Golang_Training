package history_service

// HistoryService represents history interface
type HistoryService interface {
	ExtractColor() string
	SaveColor(color string)
}
