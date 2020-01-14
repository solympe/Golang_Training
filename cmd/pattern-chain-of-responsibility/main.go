package main

import (
	"fmt"

	resp "github.com/solympe/Golang_Training/pkg/pattern-chain-of-responsibility"
)

func main() {
	plane := resp.NewPlaneDeliveryManipulator(nil)
	delivery := resp.NewCourierDeliveryManipulator(plane)
	mail := resp.NewMailDeliveryManipulator(delivery)

	fmt.Println(mail.ChooseType("couRier"))
	fmt.Println(mail.ChooseType("plAne"))

	fmt.Println(mail.ChooseType("etc"))
}
