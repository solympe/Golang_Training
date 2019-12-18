package main

import (
	"fmt"

	dc "github.com/solympe/Golang_Training/pkg/pChainOfResponsibility/deliveryCourier"
	dm "github.com/solympe/Golang_Training/pkg/pChainOfResponsibility/deliveryMail"
	dp "github.com/solympe/Golang_Training/pkg/pChainOfResponsibility/deliveryPlane"
)

func main() {
	plane := dp.NewDPlane(nil)
	delivery := dc.NewDCourier(plane)
	mail := dm.NewDMail(delivery)

	fmt.Println(mail.ChooseType("couRier"))
	fmt.Println(delivery.ChooseType("Mail"))
	fmt.Println(mail.ChooseType("plAne"))
	fmt.Println(mail.ChooseType("etc"))
}
