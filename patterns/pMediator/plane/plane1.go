package plane

import (
	ap "github.com/solympe/Golang_Training/patterns/pMediator/Airport"
)

type plane struct {
	mediator        ap.Airport
	departureDelay  int
}

func (p *plane) GetMediator (airport ap.Airport) {
	p.mediator = airport
}

func (p *plane) DelayFlight(message string, delay int) {
	p.mediator.Notify(message, delay)
	p.departureDelay += delay
}

func (p *plane) CancelFlight(message string) {
	p.mediator.Notify(message, p.departureDelay)
}

func NewPlane() ap.Aeroflot {
	return &plane{}
}