package DeliveryPlane

import (
	"strings"

	pp "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/delivery"
)

// DeliveryPlane is a plane handler struct
type DeliveryPlane struct {
	NextType pp.TypeOfDelivery
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
func NewDPlane(del pp.TypeOfDelivery) pp.TypeOfDelivery {
	return &DeliveryPlane{NextType: del}
}
