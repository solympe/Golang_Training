package subscribers

import (
	sn "github.com/solympe/Golang_Training/patterns/pObserver/subscription"
)

type subscriber struct {
	name       string
	newsPortal *sn.News
}

func (s *subscriber) Subscribe(portal sn.News) {
	s.newsPortal = &portal
	sn.News.AddSubscriber(portal, s, s.name)
}

func (s *subscriber) Unsubscribe() {
	s.newsPortal = nil
	sn.News.DeleteSubscriber(*s.newsPortal, s)
}

func NewSubscriber(name string) sn.Subscriber {
	return &subscriber{name: name}
}
