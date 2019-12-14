package DeliveryMail

import (
	p "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/delivery"
	"strings"
)

// mail handler struct
type DeliveryMail struct {
	NextType p.TypeOfDelivery
}

// checking type of delivery (if mail == true -> stop here)
func (d *DeliveryMail) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "mail" {
		return "client choosed mail delivery"
	} else if d.NextType != nil {
		return d.NextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// constructor for mail delivery
func NewDMail(del p.TypeOfDelivery) p.TypeOfDelivery {
	if del == nil {
		return &DeliveryMail{NextType: nil}
	}
	return &DeliveryMail{NextType: del}
}
