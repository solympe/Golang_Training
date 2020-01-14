package plane

import (
	"fmt"
	"strconv"

	"github.com/solympe/Golang_Training/pkg/pattern-mediator/aeroflot"
	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
)

// plane is a one of the colleague
type plane struct {
	mediator       airport.Airport
	departureDelay int
}

// GetMediator sets mediator of plane
func (p *plane) GetMediator(airport airport.Airport) {
	p.mediator = airport
}

// DelayFlight sends notify to mediator with delay time
func (p *plane) DelayFlight(delay int) {
	p.mediator.Notify("delay plane", delay)
}

// AddDelay add common delay time to plane
func (p *plane) AddDelay(delay int) {
	p.departureDelay += delay
}

// PrintDelay returns info about planes delay
func (p *plane) PrintDelay() {
	fmt.Println("Plane delay: " + strconv.Itoa(p.departureDelay) + " hours")
}

// NewHelicopter returns copy of new helicopter
func NewPlane() aeroflot.Aeroflot {
	return &plane{}
}
