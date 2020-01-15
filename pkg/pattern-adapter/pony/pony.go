package pony

import "fmt"

// Pony is a target interface
type Pony interface {
	RideOnPony()
}

type pony struct {
}

// RideOnPony ...
func (p *pony) RideOnPony() {
	fmt.Println("Im riding on pony")
}

// NewPony ...
func NewPony() Pony {
	return &pony{}
}
