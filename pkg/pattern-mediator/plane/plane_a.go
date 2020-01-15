package plane

import (
	"fmt"
	"strconv"

	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
)

// APlane ...
type APlane interface {
	DelayFlight(delay int)
	GetAirport(airport airport.Airport)
	AddDelay(delay int)
	PrintDelay()
}

type planeA struct {
	mediator       airport.Airport
	departureDelay int
}

// GetAirport sets mediator of plane
func (p *planeA) GetAirport(airport airport.Airport) {
	p.mediator = airport
}

// DelayFlight sends notify to mediator with delay time
func (p *planeA) DelayFlight(delay int) {
	if p.mediator == nil {
		panic("Mediator error!")
	}
	p.mediator.Notify("delay planeA", delay)
}

// AddDelay add common delay time to planeA
func (p *planeA) AddDelay(delay int) {
	p.departureDelay += delay
}

// PrintDelay returns info about plane delay
func (p *planeA) PrintDelay() {
	fmt.Println("Plane A delay: " + strconv.Itoa(p.departureDelay) + " hours")
}

// NewPlaneA returns copy of new planeA
func NewPlaneA() APlane {
	return &planeA{}
}
