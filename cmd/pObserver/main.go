package main

import (
	n "github.com/solympe/Golang_Training/pkg/pObserver/news"
	s "github.com/solympe/Golang_Training/pkg/pObserver/subscription"
)

func main() {
	newsPortal := n.NewNewsManipulator()
	subscriber := s.NewSubscriberManipulator("John")
	subscriber2 := s.NewSubscriberManipulator("Jake")

	newsPortal.AddSubscriber(subscriber)
	newsPortal.AddSubscriber(subscriber2)
	newsPortal.ShowSubscribers()

	newsPortal.DeleteSubscriber(subscriber)
	newsPortal.ShowSubscribers()

	newsPortal.Notify("Something")
}
