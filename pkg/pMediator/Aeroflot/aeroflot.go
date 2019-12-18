package aeroflot

import (
	"github.com/solympe/Golang_Training/pkg/pMediator/Airport"
)

// Aeroflot is a colleagues interface
type Aeroflot interface {
	DelayFlight(delay int)
	GetMediator(airport airport.Airport)
	AddDelay(delay int)
	PrintDelay()
}
