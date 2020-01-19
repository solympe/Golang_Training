package main

import (
	"github.com/solympe/Golang_Training/pkg/pattern-observer/news"
	"github.com/solympe/Golang_Training/pkg/pattern-observer/subscriber"
)

func main() {
	publisher := news.NewNews()
	subscriber1 := subscriber.NewSubscriber("John")
	subscriber2 := subscriber.NewSubscriber("Jake")
	subscriber3 := subscriber.NewSubscriber("Mike")

	publisher.Add(subscriber1)
	publisher.Add(subscriber2)
	publisher.Add(subscriber3)
	publisher.Delete(subscriber1)

	publisher.Show()

	publisher.Notify("Message")
}
