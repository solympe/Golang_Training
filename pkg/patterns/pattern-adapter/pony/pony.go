package pony

import "fmt"

// Pony is a target interface
type Pony interface {
	Ride()
}

type pony struct {
}

// Ride ...
func (p *pony) Ride() {
	fmt.Println("Im riding on pony")
}

// NewPony ...
func NewPony() Pony {
	return &pony{
	}
}
