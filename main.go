package main

import (
	"fmt"
	dc "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/deliveryCourier"
	dm "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/deliveryMail"
	dp "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/deliveryPlane"
)

func main() {
	plane := dp.NewDPlane(nil)
	delivery := dc.NewDCourier(plane)
	mail := dm.NewDMail(delivery)

	fmt.Println("1: ", mail.ChooseType("couRier"))
	fmt.Println("2: ", delivery.ChooseType("Mail"))
	fmt.Println("3: ", mail.ChooseType("plAne"))
	fmt.Println("4: ", mail.ChooseType("etc"))

}
