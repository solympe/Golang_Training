package airport

import "fmt"

type (
	model = string
	delay = int
)

type plane interface {
	Status() (model, delay)
	Delay(delay)
	Reset()
	Set(Airport)
}

type service interface {
	Set(Airport)
	Active()
	Status() bool
}

// Airport represents mediator methods
type Airport interface {
	Notify(plane, delay)
	SetPlane(plane)
	SetService(service)
	Status()
}

type airport struct {
	planes  map[plane]delay
	service service
}

// Notify gets messages from plane count its delay
func (a *airport) Notify(plane plane, delay int) {
	if a.errors() == true {
		return
	}
	a.planes[plane] += delay
	if a.planes[plane] >= 3 {
		a.service.Active()
		plane.Reset()
		a.planes[plane] = 0
	}
}

// Status prints information about delay
func (a *airport) Status() {
	if a.errors() == true {
		return
	}
	fmt.Println("Service:", a.service.Status())
	for plane, _ := range a.planes {
		fmt.Println(plane.Status())
	}
}

// SetPlane ...
func (a *airport) SetPlane(concretePlane plane) {
	if a.planes == nil {
		a.planes = make(map[plane]delay)
	}
	a.planes[concretePlane] = 0
}

// SetService
func (a *airport) SetService(service service) {
	a.service = service
}

func (a *airport) errors() bool {
	if a.service == nil || len(a.planes) == 0 {
		fmt.Println("Initialisation error")
		return true
	}
	return false
}

// NewAirport returns new copy of airport
func NewAirport() Airport {
	return &airport{}
}
