package DeliveryCourier

import (
	"strings"

	"github.com/solympe/Golang_Training/pkg/pattern-chain-of-responsibility/delivery"
)

// DeliveryCourier is a courier handler struct
type DeliveryCourier struct {
	NextType patternChainOfResponsibility.TypeOfDelivery
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
func NewDCourier(del patternChainOfResponsibility.TypeOfDelivery) patternChainOfResponsibility.TypeOfDelivery {
	return &DeliveryCourier{NextType: del}
}
