package factory

import (
	"log"

	"github.com/solympe/Golang_Training/pkg/pattern-factory/android"
)

// Factory ...
type Factory interface {
	CreateAndroid(typeOf string) android.Android
}

type factory struct{}

// CreateAndroid ...
func (f *factory) CreateAndroid(typeOf string) android.Android {
	var product android.Android
	switch typeOf {
	case "evil":
		product = android.NewAndroid("evil")
	case "kind":
		product = android.NewAndroid("kind")
	default:
		log.Fatal("Error! Unknown type")
	}
	return product
}

// NewFactory ...
func NewFactory() Factory {
	return &factory{}
}
