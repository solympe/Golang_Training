package chain_of_responsibility

import "strings"

type deliveryPlane struct {
	nextType DeliveryManipulator
}

// ChooseType checking type of delivery (if plane == true -> stop here)
func (d *deliveryPlane) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "plane" {
		return "client choosed plane delivery"
	} else if d.nextType != nil {
		return d.nextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewPlaneDeliveryManipulator is aconstructor for plane delivery
func NewPlaneDeliveryManipulator(del DeliveryManipulator) DeliveryManipulator {
	return &deliveryPlane{nextType: del}
}
