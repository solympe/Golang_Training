package DeliveryMail

import (
	"strings"

	"github.com/solympe/Golang_Training/pkg/pattern-chain-of-responsibility/delivery"
)

// DeliveryMail is a handler struct
type DeliveryMail struct {
	NextType patternChainOfResponsibility.TypeOfDelivery
}

// ChooseType checking type of delivery (if mail == true -> stop here)
func (d *DeliveryMail) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "mail" {
		return "client choosed mail delivery"
	} else if d.NextType != nil {
		return d.NextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewDMail is a constructor for mail delivery
func NewDMail(del patternChainOfResponsibility.TypeOfDelivery) patternChainOfResponsibility.TypeOfDelivery {
	return &DeliveryMail{NextType: del}
}
