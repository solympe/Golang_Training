package subscription

// SubscriberManipulator represents subscribers interface
type SubscriberManipulator interface {
	GetName() string
	GetNotify(message string)
}
