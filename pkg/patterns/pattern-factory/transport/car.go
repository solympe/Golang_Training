package transport

type carType = string

// Car ...
type Car interface {
	Get() carType
}

type car struct {
	carType string
}

func (a *car) Get() carType {
	return a.carType
}

// NewCar ...
func NewCar(carType string) Car {
	return &car{
		carType: carType,
	}
}
