package main

import (
	"github.com/solympe/Golang_Training/pkg/pattern-observer/news"
	"github.com/solympe/Golang_Training/pkg/pattern-observer/subscriber"
)

func main() {
	newsPortal := news.NewNews()
	subscriber1 := subscriber.NewSubscriber("John")
	subscriber2 := subscriber.NewSubscriber("Jake")

	newsPortal.AddSubscriber(subscriber1)
	newsPortal.AddSubscriber(subscriber2)

	newsPortal.DeleteSubscriber(subscriber1)

	newsPortal.ShowSubscribers()

	newsPortal.Notify("Message")
}
