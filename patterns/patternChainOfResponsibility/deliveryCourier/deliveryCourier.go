package DeliveryCourier

import (
	"strings"
	p "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/delivery"
)

// courier handler struct
type DeliveryCourier struct {
	NextType p.TypeOfDelivery
}

// checking type of delivery (if courier == true -> stop here)
func (d *DeliveryCourier) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "courier" {
		return "client choosed courier delivery"
	} else if d.NextType != nil {
		return d.NextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// constructor for courier delivery
func NewDCourier(del p.TypeOfDelivery) p.TypeOfDelivery {
	if del == nil {
		return &DeliveryCourier{NextType: nil}
	}
	return &DeliveryCourier{NextType: del}

}
