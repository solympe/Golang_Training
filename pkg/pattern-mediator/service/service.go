package service

import "github.com/solympe/Golang_Training/pkg/pattern-mediator/airport"

type (
	status          = bool
	concreteAirport = airport.Airport
)

// Service ...
type Service interface {
	Set(airport concreteAirport)
	Active()
	Status() status
}

type service struct {
	airport concreteAirport
	active  bool
}

// Active ...
func (s *service) Active() {
	s.active = true
}

// Set ...
func (s *service) Set(airport concreteAirport) {
	s.airport = airport
	s.airport.SetService(s)
}

// Status ...
func (s *service) Status() status {
	return s.active
}

// NewService returns copy of new service
func NewService() Service {
	return &service{}
}
