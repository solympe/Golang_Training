package delivery

// Delivery ...
type Delivery interface {
	ChooseType(chosen string) string
}
