package DeliveryPlane

import (
	"strings"

	"github.com/solympe/Golang_Training/pkg/pChainOfResponsibility/delivery"
)

// DeliveryPlane is a plane handler struct
type DeliveryPlane struct {
	NextType patternChainOfResponsibility.TypeOfDelivery
}

// ChooseType checking type of delivery (if plane == true -> stop here)
func (d *DeliveryPlane) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "plane" {
		return "client choosed plane delivery"
	} else if d.NextType != nil {
		return d.NextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewDPlane is aconstructor for plane delivery
func NewDPlane(del patternChainOfResponsibility.TypeOfDelivery) patternChainOfResponsibility.TypeOfDelivery {
	return &DeliveryPlane{NextType: del}
}
