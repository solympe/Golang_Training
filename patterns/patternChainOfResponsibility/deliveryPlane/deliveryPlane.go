package DeliveryPlane

import "strings"

// 3rd handler struct
type DeliveryPlane struct {
}

//checking type of delivery (if plane == true -> stop here) if this type is false -> return "error"
func (d *DeliveryPlane) ChooseType(chosen string) (response string) {
	if strings.ToLower(chosen) == "plane" {
		response = "client choosed plane delivery"
	} else if len(chosen) != 0 {
		response = "Delivery type error!"
	}
	return response
}
