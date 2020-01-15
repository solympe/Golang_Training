package rider

import p "github.com/solympe/Golang_Training/pkg/pattern-adapter/pony"

// Rider is a client interface
type Rider interface {
	Ride(horse p.Pony)
}

type rider struct {
}

// Ride ...
func (r *rider) Ride(horse p.Pony) {
	horse.RideOnPony()
}

// NewRider ...
func NewRider() Rider {
	return &rider{}
}
