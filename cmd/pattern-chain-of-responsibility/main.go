package main

import (
	"fmt"

	delivery "github.com/solympe/Golang_Training/pkg/pattern-chain-of-responsibility"
)

func main() {
	plane := delivery.NewPlane(nil)
	courier := delivery.NewCourier(plane)
	mail := delivery.NewMail(courier)

	result := mail.ChooseDelivery("plane")

	fmt.Println(result)
}
