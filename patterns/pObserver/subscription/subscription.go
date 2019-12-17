package subscription

// Subscriber represents subscribers interface
type Subscriber interface {
	GetName() string
	GetNotify(message string)
}
