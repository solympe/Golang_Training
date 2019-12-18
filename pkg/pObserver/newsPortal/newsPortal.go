package newsPortal

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/pObserver/news"
	"github.com/solympe/Golang_Training/pkg/pObserver/subscription"
)

// newsPortal
type newsPortal struct {
	subscribers map[subscription.Subscriber]string
}

// AddSubscriber writes new subscriber to map
func (n *newsPortal) AddSubscriber(subscriber subscription.Subscriber) {
	name := subscriber.GetName()
	n.subscribers[subscriber] = name
}

// DeleteSubscriber deletes subscriber from map
func (n *newsPortal) DeleteSubscriber(subscriber subscription.Subscriber) {
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
		subscription.Subscriber.GetNotify(key, message)
	}
}

// NewPortal - returns copy of news portal
func NewPortal() news.News {
	return &newsPortal{map[subscription.Subscriber]string{}}
}
