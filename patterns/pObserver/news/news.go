package news

import (
	s "github.com/solympe/Golang_Training/patterns/pObserver/subscription"
)

// News represents news portal interface
type News interface {
	AddSubscriber(subscriber s.Subscriber)
	DeleteSubscriber(subscriber s.Subscriber)
	ShowSubscribers()
	Notify(message string)
}
