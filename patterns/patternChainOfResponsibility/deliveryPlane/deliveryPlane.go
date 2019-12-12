package deliveryPlane

import (
	"strings"
)

type deliveryPlane struct {
}

func (p *deliveryPlane) ChooseType(chosen string) (response string) {
	if strings.ToLower(chosen) == "courier" {
		response = "client choosed courier delivery"
	} else if len(chosen) != 0 {
		response = "Error!"
	}
	return response
}

func NewDelPlane() *deliveryPlane {
	return &deliveryPlane{}
}