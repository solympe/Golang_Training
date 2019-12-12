package deliveryMail

import (
	p "github.com/solympe/Golang_Training/patterns/patternChainOfResponsibility/delivery"
	"strings"
)

type deliveryMail struct {
	nextType p.TypeOfDelivery
}

func (m *deliveryMail) ChooseType(chosen string) (response string) {
	if strings.ToLower(chosen) == "courier" {
		response = "client choosed courier delivery"
	} else if len(chosen) != 0 {
		response = m.nextType.ChooseDelivery(chosen)
	}
	return response
}

func NewDelMail(next *p.TypeOfDelivery) *deliveryMail {
	return &deliveryMail{*next}
}
