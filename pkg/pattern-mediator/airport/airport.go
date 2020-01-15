package airport

import "fmt"

type planeManager interface {
	DelayFlight(delay int)
	GetAirport(airport Airport)
	AddDelay(delay int)
	PrintDelay()
}

// Airport represents mediator methods
type Airport interface {
	Notify(message string, delay int)
	ShowStatistic()
}

type airport struct {
	planeA planeManager
	planeB planeManager
}

// Notify gets messages from plane_manager and delays departure
func (a *airport) Notify(message string, delay int) {
	if message == "delay planeA" {
		planeManager.AddDelay(a.planeA, delay)
		planeManager.AddDelay(a.planeB, delay+2)

		fmt.Println("The Plane A is delayed on", delay, ", plane B will be delayed on ", delay+2, "hours")
	}
	if message == "delay planeB" {
		planeManager.AddDelay(a.planeB, delay)
		fmt.Println("The Plane B is delayed on", delay, "hours")
	}
}

// ShowStatistic prints information about delay
func (a *airport) ShowStatistic() {
	fmt.Println("STATUS")
	planeManager.PrintDelay(a.planeA)
	planeManager.PrintDelay(a.planeB)
}

// NewAirport returns new copy of airport
func NewAirport(plane planeManager, helicopter planeManager) Airport {
	return &airport{plane, helicopter}
}
