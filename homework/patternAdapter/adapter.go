package main

import "fmt"

type children struct {}
type pony struct {}
type horse struct {}
type saddle struct {
	horseType *horse
}

func (c *children)ride(horse horsePark) {
	horse.rideOnPony()
}

type horsePark interface {
	rideOnPony()
}

func (s *saddle)rideOnPony() {
	s.horseType.rideOnHorse()
}


func (p *pony) rideOnPony () {
	fmt.Println("Im riding on pony")
}

func (h *horse) rideOnHorse() {
	fmt.Println("Im riding on horse")
}

func main() {

	Boris := &children{}
	littlePony := &pony{}
	bigHorse := &horse{}
	childSaddle := &saddle {
		horseType: bigHorse,
	}

	Boris.ride(littlePony)
	Boris.ride(childSaddle)

}