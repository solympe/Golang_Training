package helicopter

import (
	ap "github.com/solympe/Golang_Training/patterns/pMediator/Airport"
)

type helicopter struct {
	mediator       ap.Airport
	departureDelay int
}

func (h *helicopter) GetMediator(airport ap.Airport) {
	h.mediator = airport
}

func (h *helicopter) DelayFlight(message string, delay int) {
	h.mediator.Notify(message, delay)
	h.departureDelay += delay
}

func (h *helicopter) CancelFlight(message string) {
	h.mediator.Notify( message, h.departureDelay)
}

func NewHelicopter() ap.Aeroflot {
	return &helicopter{}
}
