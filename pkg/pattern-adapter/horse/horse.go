package horse

import "fmt"

// HorseRider is an adaptee interface
type HorseRider interface {
	RideOnHorse()
}

type horse struct{
}

// RideOnHorse ...
func (h *horse) RideOnHorse() {
	fmt.Println("Im riding on horse")
}

// NewHorse ...
func NewHorse() HorseRider {
	return &horse{}
}
