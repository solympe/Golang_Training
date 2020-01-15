package delivery

// Delivery ...
type Delivery interface {
	ChooseDelivery(chosen string) string
}
