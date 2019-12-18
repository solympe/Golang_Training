package pMediator

import (
	af "github.com/solympe/Golang_Training/patterns/pMediator/Aeroflot"
	ap "github.com/solympe/Golang_Training/patterns/pMediator/Airport"
)

// airport is a mediator struct
type airport struct {
	plane      *af.Aeroflot
	helicopter *af.Aeroflot
}

// Notify gets messages from aeroflot and delays departure
func (a *airport) Notify(message string, delay int) {
	if message == "delay plane" {
		af.Aeroflot.AddDelay(*a.plane, delay)
		af.Aeroflot.AddDelay(*a.helicopter, delay+2)
	}
	if message == "delay helicopter" {
		af.Aeroflot.AddDelay(*a.helicopter, delay)
		af.Aeroflot.AddDelay(*a.plane, delay+1)
	}
}

// ShowStatistic prints information about delay
func (a *airport) ShowStatistic() {
	af.Aeroflot.PrintDelay(*a.helicopter)
	af.Aeroflot.PrintDelay(*a.plane)
}

// NewAirport returns new copy of Airport
func NewAirport(plane *af.Aeroflot, helicopter *af.Aeroflot) ap.Airport {
	return &airport{plane, helicopter}
}
