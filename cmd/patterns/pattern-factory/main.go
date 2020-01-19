package main

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/patterns/pattern-factory/factory"
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-factory/transport"
)

func main() {
	transportFactory := factory.NewFactory(transport.NewCar, transport.NewTruck)

	car, _ := transportFactory.Create("car")
	truck, _ := transportFactory.Create("truck")
	fmt.Println(car.Get())
	fmt.Println(truck.Get())

	_, err := transportFactory.Create("something")
	if err != nil {
		fmt.Println(err)
	}
}
