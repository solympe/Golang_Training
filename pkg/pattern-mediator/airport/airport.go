package airport

import "fmt"

type plane interface {
	Delay(delay int)
	Set(airport Airport)
	Status()
}

// Airport represents mediator methods
type Airport interface {
	Notify(plane plane, delay int)
	Status()
	Set(plane plane)
}

type airport struct {
	planes map[plane]int
}

func (a *airport) Set(newPlane plane) {
	if a.planes == nil {
		a.planes = make(map[plane]int)
	}
	a.planes[newPlane] = 0
}

// Notify gets messages from plane_manager and delays departure
func (a *airport) Notify(plane plane, delay int) {
	a.planes[plane] += delay
}

// Status prints information about delay
func (a *airport) Status() {
	fmt.Println("STATUS")
}

// NewAirport returns new copy of airport
func NewAirport() Airport {
	return &airport{}
}
