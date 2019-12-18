package news

import (
	"github.com/solympe/Golang_Training/pkg/pObserver/subscription"
)

// News represents news portal interface
type News interface {
	AddSubscriber(subscriber subscription.Subscriber)
	DeleteSubscriber(subscriber subscription.Subscriber)
	ShowSubscribers()
	Notify(message string)
}
