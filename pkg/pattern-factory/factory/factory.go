package factory

import (
	"log"

	andr "github.com/solympe/Golang_Training/pkg/pattern-factory/android"
)

// FactoryCreator ...
type FactoryCreator interface {
	CreateAndroid(typeOf string) andr.AndroidManipulator
}

type factory struct{}

// CreateAndroid ...
func (f *factory) CreateAndroid(typeOf string) andr.AndroidManipulator {
	var product andr.AndroidManipulator
	switch typeOf {
	case "evil":
		product = andr.NewEvilAndroidManipulator()
	case "kind":
		product = andr.NewKindAndroidManipulator()
	default:
		log.Fatal("Error! Unknown type")
	}
	return product
}

// NewFactoryCreator ...
func NewFactoryCreator() FactoryCreator {
	return &factory{}
}
