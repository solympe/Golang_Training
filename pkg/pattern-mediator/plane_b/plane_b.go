package plane_b

import (
	"fmt"
	"strconv"

	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
)

// BPlane ...
type BPlane interface {
	Delay(delay int)
	Set(airport airport.Airport)
	Status()
}

type planeB struct {
	airport airport.Airport
	delay   int
}

// Set sets airport of plane_a
func (p *planeB) Set(airport airport.Airport) {
	p.airport = airport
	p.airport.Set(p)
}

// Delay sends notify to airport with delay time
func (p *planeB) Delay(delay int) {
	if p.airport == nil {
		panic("Mediator error!")
	}
	p.airport.Notify(p,delay)
}

// Status returns info about helicopters delay
func (p *planeB) Status() {
	fmt.Println("Plane B delay: " + strconv.Itoa(p.delay) + " hours")
}

// NewPlaneB returns copy of new planeB
func NewPlaneB() BPlane {
	return &planeB{}
}
