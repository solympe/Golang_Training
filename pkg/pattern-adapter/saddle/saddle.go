package saddle

import h "github.com/solympe/Golang_Training/pkg/pattern-adapter/horse"

type saddle struct {
	horseType h.HorseRider
}

// RideOnPony ...
func (s *saddle) RideOnPony() {
	s.horseType.RideOnHorse()
}

// NewSaddleRider ...
func NewSaddleRider(horse h.HorseRider) SaddleRider {
	return &saddle{horseType: horse}
}
