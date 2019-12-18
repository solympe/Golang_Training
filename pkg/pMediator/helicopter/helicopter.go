package helicopter

import (
	"fmt"
	"strconv"

	"github.com/solympe/Golang_Training/pkg/pMediator/Aeroflot"
	"github.com/solympe/Golang_Training/pkg/pMediator/Airport"
)

// helicopter is a one of the colleague
type helicopter struct {
	mediator       airport.Airport
	departureDelay int
}

// GetMediator sets mediator of helicopter
func (h *helicopter) GetMediator(airport airport.Airport) {
	h.mediator = airport
}

// DelayFlight sends notify to mediator with delay time
func (h *helicopter) DelayFlight(delay int) {
	h.mediator.Notify("delay helicopter", delay)
}

// AddDelay add common delay time to helicopter
func (h *helicopter) AddDelay(delay int) {
	h.departureDelay += delay
}

// PrintDelay returns info about helicopters delay
func (h *helicopter) PrintDelay() {
	fmt.Println("Helicopter delay: " + strconv.Itoa(h.departureDelay) + " hours")
}

// NewHelicopter returns copy of new helicopter
func NewHelicopter() aeroflot.Aeroflot {
	return &helicopter{}
}
