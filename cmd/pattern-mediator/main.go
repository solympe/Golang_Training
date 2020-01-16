package main

import (
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/plane_a"
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/plane_b"
)

func main() {
	planeA := plane_a.NewPlaneA()
	planeB := plane_b.NewPlaneB()

	planeAirport := airport.NewAirport()

	planeA.Set(planeAirport)
	planeB.Set(planeAirport)

	planeA.Delay(3)
	planeB.Delay(2)

	planeAirport.Status()
}
