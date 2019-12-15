package main

import (
	pm "github.com/solympe/Golang_Training/patterns/pMediator"
	h "github.com/solympe/Golang_Training/patterns/pMediator/helicopter"
	p "github.com/solympe/Golang_Training/patterns/pMediator/plane"

)

func main() {

	helicopter := h.NewHelicopter()
	plane := p.NewPlane()
	airport := pm.NewAirport(&plane, &helicopter)

	plane.GetMediator(airport)
	helicopter.GetMediator(airport)

	plane.DelayFlight("delay plane", 1)
	helicopter.DelayFlight("delay helicopter", 1)


}
