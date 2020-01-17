package news

import "fmt"

type (
	name    = string
	message = string
)

type subscriber interface {
	Get() name
	Report(message)
}

// News ...
type News interface {
	Add(subscriber subscriber)
	Delete(subscriber subscriber)
	Show()
	Notify(message)
}

type news struct {
	subscribers map[subscriber]name
}

// Notify - sends a message to all subscribers
func (n *news) Notify(message string) {
	for concreteSubscriber, _ := range n.subscribers {
		subscriber.Report(concreteSubscriber, message)
	}
}

// Add writes new subscriber to map
func (n *news) Add(subscriber subscriber) {
	name := subscriber.Get()
	n.subscribers[subscriber] = name
}

// Delete deletes subscriber from map
func (n *news) Delete(subscriber subscriber) {
	delete(n.subscribers, subscriber)
}

// Show show actual subscribers
func (n *news) Show() {
	fmt.Println("Subscribers now: ")
	if len(n.subscribers) == 0 {
		fmt.Println("empty")
		return
	}
	for _, val := range n.subscribers {
		fmt.Println(val + " ")
	}
}

// NewNews ...
func NewNews() News {
	return &news{map[subscriber]name{}}
}
