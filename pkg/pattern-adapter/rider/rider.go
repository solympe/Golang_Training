package rider

import p "github.com/solympe/Golang_Training/pkg/pattern-adapter/pony"

type rider struct{}

// Ride ...
func (r *rider) Ride(horse p.PonyRider) {
	horse.RideOnPony()
}

// NewRiderManipulator ...
func NewRiderManipulator() RiderManipulator {
	return &rider{}
}
