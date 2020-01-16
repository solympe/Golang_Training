package factory

import (
	"log"

	androids "github.com/solympe/Golang_Training/pkg/pattern-factory/android"
)

type android interface {
	GetType() string
}

// Factory ...
type Factory interface {
	Create(typeOf string) android
}

type factory struct {
}

// Create ...
func (f *factory) Create(typeOf string) android {
	var product android
	switch typeOf {
	case "evil":
		product = androids.NewAndroid("evil")
	case "kind":
		product = androids.NewAndroid("kind")
	default:
		log.Fatal("Error! Unknown type")
	}
	return product
}

// NewFactory ...
func NewFactory() Factory {
	return &factory{
	}
}
