package pMediator

import (
	"github.com/solympe/Golang_Training/pkg/pMediator/Aeroflot"
	airport2 "github.com/solympe/Golang_Training/pkg/pMediator/Airport"
)

// airport is a mediator struct
type airport struct {
	plane      *aeroflot.Aeroflot
	helicopter *aeroflot.Aeroflot
}

// Notify gets messages from aeroflot and delays departure
func (a *airport) Notify(message string, delay int) {
	if message == "delay plane" {
		aeroflot.Aeroflot.AddDelay(*a.plane, delay)
		aeroflot.Aeroflot.AddDelay(*a.helicopter, delay+2)
	}
	if message == "delay helicopter" {
		aeroflot.Aeroflot.AddDelay(*a.helicopter, delay)
		aeroflot.Aeroflot.AddDelay(*a.plane, delay+1)
	}
}

// ShowStatistic prints information about delay
func (a *airport) ShowStatistic() {
	aeroflot.Aeroflot.PrintDelay(*a.helicopter)
	aeroflot.Aeroflot.PrintDelay(*a.plane)
}

// NewAirport returns new copy of Airport
func NewAirport(plane *aeroflot.Aeroflot, helicopter *aeroflot.Aeroflot) airport2.Airport {
	return &airport{plane, helicopter}
}
