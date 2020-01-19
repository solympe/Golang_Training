package terminator

type (
	deliveryType = string
	response     = string
)

// Terminator ...
type Terminator interface {
	Choose(chosen deliveryType) response
}

type terminator struct {
}

// Choose ...
func (d *terminator) Choose(chosen deliveryType) response {
	return "Delivery type error"
}

// NewTerminator ...
func NewTerminator() Terminator {
	return &terminator{}
}
