package delivery

import "strings"

type mail struct {
	nextType Delivery
}

// ChooseType checking type of delivery (if mail == true -> stop here)
func (d *mail) ChooseType(chosen string) string {
	if strings.ToLower(chosen) == "mail" {
		return "client choosed mail delivery"
	} else if d.nextType != nil {
		return d.nextType.ChooseType(chosen)
	}
	return "Delivery type error"
}

// NewMail is a constructor for mail delivery
func NewMail(del Delivery) Delivery {
	return &mail{nextType: del}
}
