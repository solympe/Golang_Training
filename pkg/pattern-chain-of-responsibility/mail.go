package delivery

import "strings"

type mail struct {
	next Delivery
}

// Choose checking type of delivery (if mail == true -> stop here)
func (d *mail) Choose(chosen deliveryType) response {
	if strings.ToLower(chosen) == "mail" {
		return "client choosed mail delivery"
	} else if d.next != nil {
		return d.next.Choose(chosen)
	}
	return "Delivery type error"
}

// NewMail is a constructor for mail delivery
func NewMail(del Delivery) Delivery {
	return &mail{next: del}
}
