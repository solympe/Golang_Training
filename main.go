package main

import (
	np "github.com/solympe/Golang_Training/patterns/pObserver/newsPortal"
	s "github.com/solympe/Golang_Training/patterns/pObserver/subscribers"
)

func main() {

	newsPortal := np.NewPortal()
	subscriber := s.NewSubscriber("John")
	subscriber2 := s.NewSubscriber("Jake")

	subscriber.Subscribe(newsPortal)
	subscriber2.Subscribe(newsPortal)

	newsPortal.ShowSubscribers()

	subscriber.Unsubscribe()

}
