package main

import (
	np "github.com/solympe/Golang_Training/patterns/pObserver/newsPortal"
	s "github.com/solympe/Golang_Training/patterns/pObserver/subscribers"
)

func main() {

	newsPortal := np.NewPortal()
	subscriber := s.NewSubscriber("John")
	subscriber2 := s.NewSubscriber("Jake")

	newsPortal.AddSubscriber(subscriber)
	newsPortal.AddSubscriber(subscriber2)

	newsPortal.ShowSubscribers()

	newsPortal.DeleteSubscriber(subscriber)

	newsPortal.ShowSubscribers()

	newsPortal.Notify("Something")
}
