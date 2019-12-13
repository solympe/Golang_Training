package DeliveryMail

import (
	p "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/delivery"
	"strings"
)

// 2nd handler struct
type DeliveryMail struct {
	NextType p.TypeOfDelivery
}

// checking type of delivery (if mail == true -> stop here)
func (d *DeliveryMail) ChooseType(chosen string) (response string) {
	if strings.ToLower(chosen) == "mail" {
		response = "client choosed mail delivery"
	} else if len(chosen) != 0 {
		response = d.NextType.ChooseType(chosen)
	}
	return response
}
