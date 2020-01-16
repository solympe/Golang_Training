package airport

import "fmt"

type plane interface {
	DelayFlight(delay int)
	SetAirport(airport Airport)
	AddDelay(delay int)
	PrintDelay()
}

// Airport represents mediator methods
type Airport interface {
	Notify(message string, delay int)
	ShowCommonDelay()
}

type airport struct {
	planeA plane
	planeB plane
}

// Notify gets messages from plane_manager and delays departure
func (a *airport) Notify(message string, delay int) {
	if message == "delay planeA" {
		plane.AddDelay(a.planeA, delay)
		plane.AddDelay(a.planeB, delay+2)

		fmt.Println("The Plane A is delayed on", delay, ", plane_a B will be delayed on ", delay+2, "hours")
	}
	if message == "delay planeB" {
		plane.AddDelay(a.planeB, delay)
		fmt.Println("The Plane B is delayed on", delay, "hours")
	}
}

// ShowCommonDelay prints information about delay
func (a *airport) ShowCommonDelay() {
	fmt.Println("STATUS")
	plane.PrintDelay(a.planeA)
	plane.PrintDelay(a.planeB)
}

// NewAirport returns new copy of airport
func NewAirport(planeA plane, planeB plane) Airport {
	return &airport{planeA, planeB}
}
