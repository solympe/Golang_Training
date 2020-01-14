package planes

import (
	"fmt"
	"strconv"

	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport_manager"
)

type planeManagerB interface {
	DelayFlight(delay int)
	GetMediator(airport airport_manager.AirportManager)
	AddDelay(delay int)
	PrintDelay()
}

type planeB struct {
	mediator       airport_manager.AirportManager
	departureDelay int
}

// GetMediator sets mediator of plane_b
func (h *planeB) GetMediator(airport airport_manager.AirportManager) {
	h.mediator = airport
}

// DelayFlight sends notify to mediator with delay time
func (h *planeB) DelayFlight(delay int) {
	h.mediator.Notify("delay planeB", delay)
}

// AddDelay add common delay time to plane_b
func (h *planeB) AddDelay(delay int) {
	h.departureDelay += delay
}

// PrintDelay returns info about helicopters delay
func (h *planeB) PrintDelay() {
	fmt.Println("Plane B delay: " + strconv.Itoa(h.departureDelay) + " hours")
}

// NewPlaneManagerB returns copy of new plane_b
func NewPlaneManagerB() planeManagerB {
	return &planeB{}
}
