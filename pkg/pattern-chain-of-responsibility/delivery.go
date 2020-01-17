package delivery

type (
	deliveryType = string
	response     = string
)

// Delivery ...
type Delivery interface {
	Choose(chosen deliveryType) response
}
