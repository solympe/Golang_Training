package main

import (
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/plane"
)

func main() {
	planeA := plane.NewPlaneA()
	planeB := plane.NewPlaneB()

	planeAirport := airport.NewAirport(planeA, planeB)

	planeA.GetMediator(planeAirport)
	planeB.GetMediator(planeAirport)

	planeA.DelayFlight(3)
	planeB.DelayFlight(2)

	planeAirport.ShowStatistic()
}
