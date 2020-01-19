package mail

import "strings"

type (
	deliveryType = string
	response     = string
)

//Mail ...
type Mail interface {
	Choose(chosen deliveryType) response
}

type mail struct {
	next Mail
}

// Choose checking type of Mail (if mail == true -> stop here)
func (d *mail) Choose(chosen deliveryType) response {
	if strings.ToLower(chosen) == "mail" {
		return "client choosed mail delivery"
	}
	return d.next.Choose(chosen)
}

// NewMail ...
func NewMail(del Mail) Mail {
	return &mail{next: del}
}
