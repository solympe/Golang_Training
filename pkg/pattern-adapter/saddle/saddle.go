package saddle

import h "github.com/solympe/Golang_Training/pkg/pattern-adapter/horse"

type saddle struct {
	horseType h.HorseRider
}

func (s *saddle) RideOnPony() {
	s.horseType.RideOnHorse()
}

func NewSaddleRider(horse h.HorseRider) SaddleRider {
	return &saddle{horseType: horse}
}
