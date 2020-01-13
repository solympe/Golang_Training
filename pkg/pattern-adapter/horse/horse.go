package horse

import "fmt"

type horse struct{}

// RideOnHorse ...
func (h *horse) RideOnHorse() {
	fmt.Println("Im riding on horse")
}

// NewHorseRider ...
func NewHorseRider() HorseRider {
	return &horse{}
}
