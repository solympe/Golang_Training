package news

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/pattern-observer/subscriber"
)

// News ...
type News interface {
	AddSubscriber(subscriber subscriber.Subscriber)
	DeleteSubscriber(subscriber subscriber.Subscriber)
	ShowSubscribers()
	Notify(message string)
}

type news struct {
	subscribers map[subscriber.Subscriber]string
}

// AddSubscriber writes new subscriber to map
func (n *news) AddSubscriber(subscriber subscriber.Subscriber) {
	name := subscriber.GetName()
	n.subscribers[subscriber] = name
}

// DeleteSubscriber deletes subscriber from map
func (n *news) DeleteSubscriber(subscriber subscriber.Subscriber) {
	delete(n.subscribers, subscriber)
}

// ShowSubscribers show actual subscribers
func (n *news) ShowSubscribers() {
	fmt.Print("Subscribers now: ")
	if len(n.subscribers) == 0 {
		fmt.Println("empty")
		return
	}
	for _, val := range n.subscribers {
		fmt.Println(val + " ")
	}
}

// Notify - sends a message to all subscribers
func (n *news) Notify(message string) {
	for key, _ := range n.subscribers {
		subscriber.Subscriber.GetNotify(key, message)
	}
}

// NewNews ...
func NewNews() News {
	return &news{map[subscriber.Subscriber]string{}}
}
