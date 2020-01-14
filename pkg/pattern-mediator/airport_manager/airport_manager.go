package airport_manager

// AirportManager represents mediator methods
type AirportManager interface {
	Notify(message string, delay int)
	ShowStatistic()
}
