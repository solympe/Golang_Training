package subscriber

import "fmt"

type (
	name    = string
	message = string
)

// Subscriber ...
type Subscriber interface {
	Get() name
	Report(message)
}

type subscriber struct {
	name string
}

// Get return subscriber`s name
func (s *subscriber) Get() name {
	return s.name
}

// Report ...
func (s *subscriber) Report(message string) {
	fmt.Println(s.name + " received: " + message)
}

// NewSubscriber returns new copy of subscriber
func NewSubscriber(name string) Subscriber {
	return &subscriber{name: name}
}
