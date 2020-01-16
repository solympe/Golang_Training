package rider

type pony interface {
	Ride()
}

// Rider is a client interface
type Rider interface {
	Ride(horse pony)
}

type rider struct {
}

// Ride ...
func (r *rider) Ride(horse pony) {
	horse.Ride()
}

// NewRider ...
func NewRider() Rider {
	return &rider{
	}
}
