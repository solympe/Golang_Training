package main

import (
	"fmt"
	p "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/delivery"
	dc "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/deliveryCourier"
	dm "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/deliveryMail"
	dp "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/deliveryPlane"
)

func main() {

	chosenDelivery := "courier"

	var plane = dp.NewDelPlane()
	var mail = dm.NewDelMail(plane)
	var courier = dc.NewDelCourier(mail)



	fmt.Println(chosenDelivery)
}
