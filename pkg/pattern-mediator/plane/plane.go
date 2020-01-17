package plane

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"
)

type (
	model           = string
	delay           = int
	concreteAirport = airport.Airport
)

// Plane ...
type Plane interface {
	Status() (model, delay)
	Delay(delay)
	Reset()
	Set(airport concreteAirport)
}

type plane struct {
	model   string
	delay   int
	airport concreteAirport
}

// Set sets airport
func (p *plane) Set(airport concreteAirport) {
	p.airport = airport
	p.airport.SetPlane(p)
}

// Delay ...
func (p *plane) Delay(delay int) {
	if p.airport == nil {
		fmt.Println("Airport init error for", p.model)
		return
	}
	p.delay += delay
	p.airport.Notify(p, delay)
}

// Reset ...
func (p *plane) Reset() {
	p.delay = 0
}

// Status ...
func (p *plane) Status() (model, delay) {
	return p.model, p.delay
}

// NewPlane returns copy of new plane
func NewPlane(model string) Plane {
	return &plane{
		model: model,
	}
}
