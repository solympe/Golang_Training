package planes

import (
	"fmt"
	"strconv"

	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport_manager"
)

type planeManagerA interface {
	DelayFlight(delay int)
	GetMediator(airport airport_manager.AirportManager)
	AddDelay(delay int)
	PrintDelay()
}

type planeA struct {
	mediator       airport_manager.AirportManager
	departureDelay int
}

// GetMediator sets mediator of planes
func (p *planeA) GetMediator(airport airport_manager.AirportManager) {
	p.mediator = airport
}

// DelayFlight sends notify to mediator with delay time
func (p *planeA) DelayFlight(delay int) {
	p.mediator.Notify("delay planeA", delay)
}

// AddDelay add common delay time to planeA
func (p *planeA) AddDelay(delay int) {
	p.departureDelay += delay
}

// PrintDelay returns info about planes delay
func (p *planeA) PrintDelay() {
	fmt.Println("Plane A delay: " + strconv.Itoa(p.departureDelay) + " hours")
}

// NewPlaneManagerA returns copy of new planeA
func NewPlaneManagerA() planeManagerA {
	return &planeA{}
}
