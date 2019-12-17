package newsPortal

import (
	"fmt"

	n "github.com/solympe/Golang_Training/patterns/pObserver/news"
	s "github.com/solympe/Golang_Training/patterns/pObserver/subscription"
)

// newsPortal
type newsPortal struct {
	subscribers map[s.Subscriber]string
}

// AddSubscriber writes new subscriber to map
func (n *newsPortal) AddSubscriber(subscriber s.Subscriber) {
	name := subscriber.GetName()
	n.subscribers[subscriber] = name
}

// DeleteSubscriber deletes subscriber from map
func (n *newsPortal) DeleteSubscriber(subscriber s.Subscriber) {
	delete(n.subscribers, subscriber)
}

// ShowSubscribers show actual subscribers
func (n *newsPortal) ShowSubscribers() {
	fmt.Print("Subscribers now: ")
	for _, val := range n.subscribers {
		fmt.Print(val + " ")
	}
	fmt.Println()
}

// Notify - sends a message to all subscribers
func (n *newsPortal) Notify(message string) {
	for key, _ := range n.subscribers {
		s.Subscriber.GetNotify(key, message)
	}
}

// NewPortal - returns copy of news portal
func NewPortal() n.News {
	return &newsPortal{map[s.Subscriber]string{}}
}
