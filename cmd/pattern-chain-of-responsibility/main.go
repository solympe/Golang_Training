package main

import (
	"fmt"

	delivery "github.com/solympe/Golang_Training/pkg/pattern-chain-of-responsibility"
)

func main() {
	plane := delivery.NewPlane(nil)
	deliveryA := delivery.NewCourier(plane)
	mail := delivery.NewMail(deliveryA)

	fmt.Println(mail.ChooseType("courier"))
	fmt.Println(mail.ChooseType("plAne"))

	fmt.Println(mail.ChooseType("etc"))
}
