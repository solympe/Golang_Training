package delivery

import "strings"

type courier struct {
	next Delivery
}

// Choose checking type of delivery (if courier == true -> stop here)
func (d *courier) Choose(chosen deliveryType) response {
	if strings.ToLower(chosen) == "courier" {
		return "client choosed courier delivery"
	} else if d.next != nil {
		return d.next.Choose(chosen)
	}
	return "Delivery type error"
}

// NewCourier is a constructor for courier delivery
func NewCourier(del Delivery) Delivery {
	return &courier{next: del}
}
