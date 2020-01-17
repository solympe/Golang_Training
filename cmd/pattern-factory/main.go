package main

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/pattern-factory/factory"
)

func main() {
	transportFactory := factory.NewFactory()

	car := transportFactory.Create("car")
	truck := transportFactory.Create("truck")

	fmt.Println(car.Get())
	fmt.Println(truck.Get())
}
