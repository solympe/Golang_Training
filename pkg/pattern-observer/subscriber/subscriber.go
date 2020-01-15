package subscriber

import "fmt"

// Subscriber ...
type Subscriber interface {
	GetName() string
	GetNotify(message string)
}

type subscriber struct {
	name string
}

// GetName return subscriber`s name
func (s *subscriber) GetName() string {
	return s.name
}

// GetNotify returns recieved message
func (s *subscriber) GetNotify(message string) {
	fmt.Println("I am " + s.name + " and I received: " + message)
}

// NewSubscriber returns new copy of subscriber
func NewSubscriber(name string) Subscriber {
	return &subscriber{name: name}
}
