package saddle

// Saddle is an adapter
type Saddle interface {
	Ride()
}

type saddle struct {
	horseType Saddle
}

// Ride ...
func (s *saddle) Ride() {
	s.horseType.Ride()
}

// NewSaddle ...
func NewSaddle(horse Saddle) Saddle {
	return &saddle{horseType: horse}
}
