package chain_of_responsibility

// DeliveryManipulator ...
type DeliveryManipulator interface {
	ChooseType(chosen string) string
}
