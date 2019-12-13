package main

import (
	"fmt"
	dc "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/deliveryCourier"
)

func main() {
	delivery := dc.NewDCourier()

	fmt.Println(delivery.ChooseType("couRier"))
	fmt.Println(delivery.ChooseType("Mail"))
	fmt.Println(delivery.ChooseType("plane"))
	fmt.Println(delivery.ChooseType("something else"))

}
