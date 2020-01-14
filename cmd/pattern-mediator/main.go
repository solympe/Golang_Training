package main

import (
	airport2 "github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
	pm "github.com/solympe/Golang_Training/pkg/pattern-mediator/plane_manager"
	p "github.com/solympe/Golang_Training/pkg/pattern-mediator/planes"
)

func main() {
	helicopter := p.NewPlaneManagerB()
	plane := p.NewPlaneManagerA()
	airport := airport2.NewAirportManager((*pm.PlaneManager)(&plane), (*pm.PlaneManager)(&helicopter))

	plane.GetMediator(airport)
	helicopter.GetMediator(airport)

	plane.DelayFlight(3)
	helicopter.DelayFlight(1)

	airport.ShowStatistic()
}
