package news

import "github.com/solympe/Golang_Training/pkg/pObserver/subscription"

// News represents news portal interface
type NewsManipulator interface {
	AddSubscriber(subscriber subscription.SubscriberManipulator)
	DeleteSubscriber(subscriber subscription.SubscriberManipulator)
	ShowSubscribers()
	Notify(message string)
}
