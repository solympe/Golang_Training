package factory

import (
	"errors"

	"github.com/solympe/Golang_Training/pkg/pattern-factory/transport"
)

type (
	transportType = string
	err           = error
	carCreator    func(carType string) transport.Car
	truckCreator  func(truckType string) transport.Truck
)

type createdTransport interface {
	Get() transportType
}

// Factory ...
type Factory interface {
	Create(transportType) (createdTransport, err)
}

type factory struct {
	carCreator   carCreator
	truckCreator truckCreator
}

func (f *factory) Create(transportType string) (product createdTransport, err error) {
	switch transportType {
	case "truck":
		product = f.truckCreator("truck")
		return
	case "car":
		product = f.carCreator("car")
		return
	}
	err = errors.New("creation error")
	return
}

// NewFactory ...
func NewFactory(carCreator carCreator, truckCreator truckCreator) Factory {
	return &factory{
		carCreator:   carCreator,
		truckCreator: truckCreator,
	}
}
