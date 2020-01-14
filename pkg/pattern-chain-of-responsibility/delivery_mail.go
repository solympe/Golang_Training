package chain_of_responsibility

import "strings"

type deliveryMail struct {
	nextType DeliveryManipulator
}

// ChooseType checking type of delivery (if mail == true -> stop here)
func (d *deliveryMail) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "mail" {
		return "client choosed mail delivery"
	} else if d.nextType != nil {
		return d.nextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewMailDeliveryManipulator is a constructor for mail delivery
func NewMailDeliveryManipulator(del DeliveryManipulator) DeliveryManipulator {
	return &deliveryMail{nextType: del}
}
