package patternChainOfResponsibility

type TypeOfDelivery interface {
	ChooseDelivery(chosen string)
}
