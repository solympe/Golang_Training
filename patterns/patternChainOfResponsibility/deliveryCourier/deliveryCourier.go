package DeliveryCourier

import (
	"strings"

	p "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/delivery"
)

// DeliveryCourier is a courier handler struct
type DeliveryCourier struct {
	NextType p.TypeOfDelivery
}

// ChooseType checking type of delivery (if courier == true -> stop here)
func (d *DeliveryCourier) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "courier" {
		return "client choosed courier delivery"
	} else if d.NextType != nil {
		return d.NextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewDCourier is a constructor for courier delivery
func NewDCourier(del p.TypeOfDelivery) p.TypeOfDelivery {
	return &DeliveryCourier{NextType: del}
}
