package news

import "github.com/solympe/Golang_Training/pkg/pattern-observer/subscription"

// NewsManipulator ...
type NewsManipulator interface {
	AddSubscriber(subscriber subscription.SubscriberManipulator)
	DeleteSubscriber(subscriber subscription.SubscriberManipulator)
	ShowSubscribers()
	Notify(message string)
}
