package rider

import p "github.com/solympe/Golang_Training/pkg/pattern-adapter/pony"

// RiderManipulator is a client interface
type RiderManipulator interface {
	Ride(horse p.PonyRider)
}
