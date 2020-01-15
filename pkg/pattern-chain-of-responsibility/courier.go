package delivery

import "strings"

type courier struct {
	nextType Delivery
}

// ChooseType checking type of delivery (if courier == true -> stop here)
func (d *courier) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "courier" {
		return "client choosed courier delivery"
	} else if d.nextType != nil {
		return d.nextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewCourier is a constructor for courier delivery
func NewCourier(del Delivery) Delivery {
	return &courier{nextType: del}
}
