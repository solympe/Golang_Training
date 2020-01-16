package main

import (
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/plane_a"
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/plane_b"
)

func main() {
	planeA := plane_a.NewPlaneA()
	planeB := plane_b.NewPlaneB()

	planeAirport := airport.NewAirport(planeA, planeB)

	planeA.SetAirport(planeAirport)
	planeB.SetAirport(planeAirport)

	planeA.DelayFlight(3)
	planeB.DelayFlight(2)

	planeAirport.ShowCommonDelay()
}
