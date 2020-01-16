package DeliveryPlane

import (
	"strings"

	"github.com/solympe/Golang_Training/pkg/pattern-chain-of-responsibility/delivery"
)

// DeliveryPlane is a plane_a handler struct
type DeliveryPlane struct {
	NextType patternChainOfResponsibility.TypeOfDelivery
}

// ChooseType checking type of delivery (if plane_a == true -> stop here)
func (d *DeliveryPlane) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "plane_a" {
		return "client choosed plane_a delivery"
	} else if d.NextType != nil {
		return d.NextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewDPlane is aconstructor for plane_a delivery
func NewDPlane(del patternChainOfResponsibility.TypeOfDelivery) patternChainOfResponsibility.TypeOfDelivery {
	return &DeliveryPlane{NextType: del}
}
