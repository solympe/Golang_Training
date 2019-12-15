package pMediator

import (
	"fmt"
	"strconv"

	ap "github.com/solympe/Golang_Training/patterns/pMediator/Airport"
)

type airport struct {
	plane      *ap.Aeroflot
	helicopter *ap.Aeroflot
}

// добавить распознавание отправителя
func (a *airport) Notify( message string, delay int) {
	if message == "delay plane" {
		fmt.Println("Plane departure will delay on " + strconv.Itoa(delay) + " hours")
	}
	if message == "delay helicopter" {
		fmt.Println("Helicopter departure will delay on " + strconv.Itoa(delay) + " hours")
	}
}

func NewAirport(plane *ap.Aeroflot, helicopter *ap.Aeroflot) ap.Airport {
	return &airport{plane, helicopter}
}
