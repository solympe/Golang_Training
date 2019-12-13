package DeliveryCourier

import (
	p "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/delivery"
	dm "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/deliveryMail"
	dp "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/deliveryPlane"
	"strings"
)

// 1st handler struct
type DeliveryCourier struct {
	NextType p.TypeOfDelivery
}

// checking type of delivery (if courier == true -> stop here)
func (d *DeliveryCourier) ChooseType(chosen string) (response string) {
	if strings.ToLower(chosen) == "courier" {
		response = "client choosed courier delivery"
	} else if len(chosen) != 0 {
		response = d.NextType.ChooseType(chosen)
	}
	return response
}

// handlers constructor
func NewDCourier() p.TypeOfDelivery {
	return &DeliveryCourier{NextType: &dm.DeliveryMail{NextType: &dp.DeliveryPlane{}}}
}