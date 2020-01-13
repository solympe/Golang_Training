package pony

import "fmt"

type pony struct{}

// RideOnPony ...
func (p *pony) RideOnPony() {
	fmt.Println("Im riding on pony")
}

// NewPonyRider ...
func NewPonyRider() PonyRider {
	return &pony{}
}
