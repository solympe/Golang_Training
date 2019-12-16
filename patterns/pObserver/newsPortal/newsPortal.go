package newsPortal

import (
	"fmt"

	s "github.com/solympe/Golang_Training/patterns/pObserver/subscription"
)

type newsPortal struct {
	subscribers map[s.Subscriber]string
}

func (n *newsPortal) AddSubscriber(subscriber s.Subscriber, name string) {
	n.subscribers[subscriber] = name
}

func (n *newsPortal) DeleteSubscriber(subscriber s.Subscriber) {
	fmt.Println("aaaaA!", subscriber)
	delete(n.subscribers, subscriber)
}

func (n *newsPortal) ShowSubscribers() {
	fmt.Println(n.subscribers)
}

func (n *newsPortal) Notify() {
	fmt.Println("ON WORKING")
}

func NewPortal() s.News {
	return &newsPortal{map[s.Subscriber]string{}}
}
