package patternChainOfResponsibility

// common interface for handlers
type TypeOfDelivery interface {
	ChooseType(chosen string) string
}
