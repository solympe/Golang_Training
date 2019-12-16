package aeroflot

import ap "github.com/solympe/Golang_Training/patterns/pMediator/Airport"

// Aeroflot is a colleagues interface
type Aeroflot interface {
	DelayFlight(delay int)
	GetMediator(airport ap.Airport)
	AddDelay(delay int)
	PrintDelay()
}
