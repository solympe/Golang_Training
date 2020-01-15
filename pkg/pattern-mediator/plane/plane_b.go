package plane

import (
	"fmt"
	"strconv"

	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
)

// BPlane ...
type BPlane interface {
	DelayFlight(delay int)
	GetMediator(airport airport.Airport)
	AddDelay(delay int)
	PrintDelay()
}

type planeB struct {
	mediator       airport.Airport
	departureDelay int
}

// GetMediator sets mediator of plane_b
func (p *planeB) GetMediator(airport airport.Airport) {
	p.mediator = airport
}

// DelayFlight sends notify to mediator with delay time
func (p *planeB) DelayFlight(delay int) {
	if p.mediator == nil {
		panic("Mediator error!")
	}
	p.mediator.Notify("delay planeB", delay)
}

// AddDelay add common delay time to planeB
func (p *planeB) AddDelay(delay int) {
	p.departureDelay += delay
}

// PrintDelay returns info about helicopters delay
func (p *planeB) PrintDelay() {
	fmt.Println("Plane B delay: " + strconv.Itoa(p.departureDelay) + " hours")
}

// NewPlaneB returns copy of new planeB
func NewPlaneB() BPlane {
	return &planeB{}
}
