package airport

type Airport interface {
	Notify(message string, delay int)
}

type Aeroflot interface {
	DelayFlight (message string, delay int)
	GetMediator (airport Airport)
	CancelFlight(message string)
}