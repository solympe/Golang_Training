package plane

import "strings"

type (
	deliveryType = string
	response     = string
)

//Plane ...
type Plane interface {
	Choose(chosen deliveryType) response
}

type plane struct {
	next Plane
}

// Choose checking type of Plane (if plane == true -> stop here)
func (d *plane) Choose(chosen deliveryType) response {
	if strings.ToLower(chosen) == "plane" {
		return "client choosed plane delivery"
	}
	return d.next.Choose(chosen)
}

// NewPlane ...
func NewPlane(del Plane) Plane {
	return &plane{next: del}
}
