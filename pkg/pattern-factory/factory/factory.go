package factory

import (
	"log"

	"github.com/solympe/Golang_Training/pkg/pattern-factory/transport"
)

type transportType = string

type createdTransport interface {
	Get() transportType
}

// Factory ...
type Factory interface {
	Create(transportType) createdTransport
}

type factory struct {
}

// Create ...
func (f *factory) Create(transportType string) (product createdTransport) {
	switch transportType {
	case "truck":
		product = transport.NewTruck()
	case "car":
		product = transport.NewCar()
	default:
		log.Fatal("Error! Unknown type of product")
	}
	return product
}

// NewFactory ...
func NewFactory() Factory {
	return &factory{}
}
