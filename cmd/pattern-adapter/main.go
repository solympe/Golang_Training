package main

import (
	"github.com/solympe/Golang_Training/pkg/pattern-adapter/horse"
	"github.com/solympe/Golang_Training/pkg/pattern-adapter/pony"
	"github.com/solympe/Golang_Training/pkg/pattern-adapter/rider"
	"github.com/solympe/Golang_Training/pkg/pattern-adapter/saddle"
)

func main() {
	horseRider := rider.NewRiderManipulator()

	targetPony := pony.NewPonyRider()
	adapteeHorse := horse.NewHorseRider()

	adapterSaddle := saddle.NewSaddleRider(adapteeHorse)

	horseRider.Ride(targetPony)
	horseRider.Ride(adapterSaddle)
}
