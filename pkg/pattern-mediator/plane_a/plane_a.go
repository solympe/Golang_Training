package plane_a

import (
	"fmt"
	"strconv"

	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
)

// APlane ...
type APlane interface {
	Delay(delay int)
	Set(airport airport.Airport)
	Status()
}

type planeA struct {
	airport airport.Airport
	delay   int
}

// Set sets airport of plane_a
func (p *planeA) Set(airport airport.Airport) {
	p.airport = airport
	p.airport.Set(p)
}

// Delay sends notify to airport with delay time
func (p *planeA) Delay(delay int) {
	if p.airport == nil {
		panic("Mediator error!")
	}
	p.delay += delay
	p.airport.Notify(p, delay)
}

// Status returns info about plane_a delay
func (p *planeA) Status() {
	fmt.Println("Plane A delay: " + strconv.Itoa(p.delay) + " hours")
}

// NewPlaneA returns copy of new planeA
func NewPlaneA() APlane {
	return &planeA{}
}
