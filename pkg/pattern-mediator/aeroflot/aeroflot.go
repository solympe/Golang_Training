package aeroflot

import (
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
)

// aeroflot is a colleagues interface
type Aeroflot interface {
	DelayFlight(delay int)
	GetMediator(airport airport.Airport)
	AddDelay(delay int)
	PrintDelay()
}
