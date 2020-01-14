package DeliveryPlane

import (
	"strings"

	"github.com/solympe/Golang_Training/pkg/pattern-chain-of-responsibility/delivery"
)

// DeliveryPlane is a planes handler struct
type DeliveryPlane struct {
	NextType patternChainOfResponsibility.TypeOfDelivery
}

// ChooseType checking type of delivery (if planes == true -> stop here)
func (d *DeliveryPlane) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "planes" {
		return "client choosed planes delivery"
	} else if d.NextType != nil {
		return d.NextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewDPlane is aconstructor for planes delivery
func NewDPlane(del patternChainOfResponsibility.TypeOfDelivery) patternChainOfResponsibility.TypeOfDelivery {
	return &DeliveryPlane{NextType: del}
}
