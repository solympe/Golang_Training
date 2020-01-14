package plane_manager

import "github.com/solympe/Golang_Training/pkg/pattern-mediator/airport_manager"

// PlaneManager ...
type PlaneManager interface {
	DelayFlight(delay int)
	GetMediator(airport airport_manager.AirportManager)
	AddDelay(delay int)
	PrintDelay()
}
