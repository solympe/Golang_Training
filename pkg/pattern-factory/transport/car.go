package transport

type carType = string

// Car ...
type Car interface {
	Get() carType
}

type car struct {
	carType string
}

// Get ...
func (a *car) Get() carType {
	return a.carType
}

// NewCar ...
func NewCar() Car {
	return &car{
		carType: "car",
	}
}
