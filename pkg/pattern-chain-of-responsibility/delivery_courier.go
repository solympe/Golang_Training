package chain_of_responsibility

import "strings"

type deliveryCourier struct {
	nextType DeliveryManipulator
}

// ChooseType checking type of delivery (if courier == true -> stop here)
func (d *deliveryCourier) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "courier" {
		return "client choosed courier delivery"
	} else if d.nextType != nil {
		return d.nextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewCourierDeliveryManipulator is a constructor for courier delivery
func NewCourierDeliveryManipulator(del DeliveryManipulator) DeliveryManipulator {
	return &deliveryCourier{nextType: del}
}
