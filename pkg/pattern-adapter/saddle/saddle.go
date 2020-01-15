package saddle

import h "github.com/solympe/Golang_Training/pkg/pattern-adapter/horse"

// Saddle is an adapter
type Saddle interface {
	RideOnPony()
}

type saddle struct {
	horseType h.HorseRider
}

// RideOnPony ...
func (s *saddle) RideOnPony() {
	s.horseType.RideOnHorse()
}

// NewSaddle ...
func NewSaddle(horse h.HorseRider) Saddle {
	return &saddle{horseType: horse}
}
