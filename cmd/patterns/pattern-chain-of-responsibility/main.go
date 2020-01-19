package main

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/patterns/pattern-chain-of-responsibility/courier"
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-chain-of-responsibility/mail"
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-chain-of-responsibility/plane"
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-chain-of-responsibility/terminator"
)

func main() {
	end := terminator.NewTerminator()

	deliveryPlane := plane.NewPlane(end)
	deliveryCourier := courier.NewCourier(deliveryPlane)
	deliveryMail := mail.NewMail(deliveryCourier)

	result := deliveryMail.Choose("mail")
	fmt.Println(result)
}
