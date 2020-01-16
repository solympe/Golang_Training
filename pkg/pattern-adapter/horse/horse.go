package horse

import "fmt"

// Horse is an adaptee interface
type Horse interface {
	Ride()
}

type horse struct {
}

// Ride ...
func (h *horse) Ride() {
	fmt.Println("Im riding on horse")
}

// NewHorse ...
func NewHorse() Horse {
	return &horse{
	}
}
