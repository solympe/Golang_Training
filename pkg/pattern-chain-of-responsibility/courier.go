package delivery

import "strings"

type courier struct {
	nextType Delivery
}

// ChooseDelivery checking type of delivery (if courier == true -> stop here)
func (d *courier) ChooseDelivery(chosen string) string {
	if strings.ToLower(chosen) == "courier" {
		return "client choosed courier delivery"
	} else if d.nextType != nil {
		return d.nextType.ChooseDelivery(chosen)
	}
	return "Delivery type error"
}

// NewCourier is a constructor for courier delivery
func NewCourier(del Delivery) Delivery {
	return &courier{nextType: del}
}
