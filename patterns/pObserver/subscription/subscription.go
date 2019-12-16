package subscription


type News interface {
	AddSubscriber (subscriber Subscriber, name string)
	DeleteSubscriber (subscriber Subscriber)
	ShowSubscribers ()
	Notify ()
}

type Subscriber interface {
	Subscribe(portal News, )
	Unsubscribe()
}

// переделать