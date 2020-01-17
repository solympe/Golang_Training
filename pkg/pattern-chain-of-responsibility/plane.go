package delivery

import "strings"

type plane struct {
	next Delivery
}

// Choose checking type of delivery (if plane == true -> stop here)
func (d *plane) Choose(chosen deliveryType) response {
	if strings.ToLower(chosen) == "plane" {
		return "client choosed plane delivery"
	} else if d.next != nil {
		return d.next.Choose(chosen)
	}
	return "Delivery type error"
}

// NewPlane is aconstructor for plane delivery
func NewPlane(del Delivery) Delivery {
	return &plane{next: del}
}
