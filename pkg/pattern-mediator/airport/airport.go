package airport

import (
	ap "github.com/solympe/Golang_Training/pkg/pattern-mediator/airport_manager"
	pm "github.com/solympe/Golang_Training/pkg/pattern-mediator/plane_manager"
)

type airport struct {
	planeA *pm.PlaneManager
	planeB *pm.PlaneManager
}

// Notify gets messages from plane_manager and delays departure
func (a *airport) Notify(message string, delay int) {
	if message == "delay planeA" {
		pm.PlaneManager.AddDelay(*a.planeA, delay)
		pm.PlaneManager.AddDelay(*a.planeB, delay+2)
	}
	if message == "delay planeB" {
		pm.PlaneManager.AddDelay(*a.planeB, delay)
	}
}

// ShowStatistic prints information about delay
func (a *airport) ShowStatistic() {
	pm.PlaneManager.PrintDelay(*a.planeB)
	pm.PlaneManager.PrintDelay(*a.planeA)
}

// NewAirport returns new copy of airport_manager
func NewAirportManager(plane *pm.PlaneManager, helicopter *pm.PlaneManager) ap.AirportManager {
	return &airport{plane, helicopter}
}
