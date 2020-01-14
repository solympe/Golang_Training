package subscription

// SubscriberManipulator ...
type SubscriberManipulator interface {
	GetName() string
	GetNotify(message string)
}
