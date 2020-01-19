package main

import (
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-adapter/horse"
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-adapter/pony"
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-adapter/rider"
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-adapter/saddle"
)

func main() {
	ponyRider := rider.NewRider()
	targetPony := pony.NewPony()

	adapteeHorse := horse.NewHorse()
	adapterSaddle := saddle.NewSaddle(adapteeHorse)

	ponyRider.Ride(targetPony)
	ponyRider.Ride(adapterSaddle)
}
