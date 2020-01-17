package courier

import "strings"

type (
	deliveryType = string
	response     = string
)

// Courier ...
type Courier interface {
	Choose(chosen deliveryType) response
}

type courier struct {
	next Courier
}

// Choose checking type of Courier (if courier == true -> stop here)
func (d *courier) Choose(chosen deliveryType) response {
	if strings.ToLower(chosen) == "courier" {
		return "client choosed courier delivery"
	}
	return d.next.Choose(chosen)
}

// NewCourier ...
func NewCourier(del Courier) Courier {
	return &courier{next: del}
}
