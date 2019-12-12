package deliveryCourier

import (
	p "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/delivery"
	"strings"
)

type deliveryCourier struct {
	nextType p.TypeOfDelivery
}

func (d *deliveryCourier) ChooseType(chosen string) (response string) {
	if strings.ToLower(chosen) == "courier" {
		response = "client choosed courier delivery"
	} else if len(chosen) != 0 {
		response = d.nextType.ChooseDelivery()
	}
	return response
}

func NewDelCourier(next *p.TypeOfDelivery) *deliveryCourier{
	return &deliveryCourier{*next}
}