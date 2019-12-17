package patternChainOfResponsibility

// TypeOfDelivery is a common interface for handlers
type TypeOfDelivery interface {
	ChooseType(chosen string) string
}
