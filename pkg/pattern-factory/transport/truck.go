package transport

type truckType = string

// Truck ...
type Truck interface {
	Get() truckType
}

type truck struct {
	truckType string
}

// Get ...
func (a *truck) Get() truckType {
	return a.truckType
}

// NewTruck ...
func NewTruck() Truck {
	return &truck{
		truckType: "truck",
	}
}
