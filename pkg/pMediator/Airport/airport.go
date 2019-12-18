package airport

// Airport represents mediator interface
type Airport interface {
	Notify(message string, delay int)
	ShowStatistic()
}
