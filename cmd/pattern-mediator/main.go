package main

import (
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/plane"
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/service"
)

func main() {
	planeA := plane.NewPlane("Plane A")
	planeB := plane.NewPlane("Plane B")
	planeService := service.NewService()
	planeAirport := airport.NewAirport()

	planeA.Set(planeAirport)
	planeB.Set(planeAirport)
	planeService.Set(planeAirport)

	planeA.Delay(3)
	planeB.Delay(2)

	planeAirport.Status()
}
