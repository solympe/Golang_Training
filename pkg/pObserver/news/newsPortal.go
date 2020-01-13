package news

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/pObserver/subscription"
)

type newsPortal struct {
	subscribers map[subscription.SubscriberManipulator]string
}

// AddSubscriber writes new subscriber to map
func (n *newsPortal) AddSubscriber(subscriber subscription.SubscriberManipulator) {
	name := subscriber.GetName()
	n.subscribers[subscriber] = name
}

// DeleteSubscriber deletes subscriber from map
func (n *newsPortal) DeleteSubscriber(subscriber subscription.SubscriberManipulator) {
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
		subscription.SubscriberManipulator.GetNotify(key, message)
	}
}

// NewPortal - returns copy of news portal
func NewPortal() News {
	return &newsPortal{map[subscription.SubscriberManipulator]string{}}
}
