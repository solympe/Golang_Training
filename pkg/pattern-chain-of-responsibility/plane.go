package delivery

import "strings"

type plane struct {
	nextType Delivery
}

// ChooseType checking type of delivery (if plane == true -> stop here)
func (d *plane) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "plane" {
		return "client choosed plane delivery"
	} else if d.nextType != nil {
		return d.nextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewPlane is aconstructor for plane delivery
func NewPlane(del Delivery) Delivery {
	return &plane{nextType: del}
}
