package news

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/pObserver/subscription"
)

type news struct {
	subscribers map[subscription.SubscriberManipulator]string
}

// AddSubscriber writes new subscriber to map
func (n *news) AddSubscriber(subscriber subscription.SubscriberManipulator) {
	name := subscriber.GetName()
	n.subscribers[subscriber] = name
}

// DeleteSubscriber deletes subscriber from map
func (n *news) DeleteSubscriber(subscriber subscription.SubscriberManipulator) {
	delete(n.subscribers, subscriber)
}

// ShowSubscribers show actual subscribers
func (n *news) ShowSubscribers() {
	fmt.Print("Subscribers now: ")
	for _, val := range n.subscribers {
		fmt.Print(val + " ")
	}
	fmt.Println()
}

// Notify - sends a message to all subscribers
func (n *news) Notify(message string) {
	for key, _ := range n.subscribers {
		subscription.SubscriberManipulator.GetNotify(key, message)
	}
}

// NewNews - returns news
func NewNews() NewsManipulator {
	return &news{map[subscription.SubscriberManipulator]string{}}
}
