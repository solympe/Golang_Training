package android

// Android ...
type Android interface {
	GetType() string
}

type android struct {
	typeOf string
}

// GetType ...
func (a *android) GetType() string {
	return a.typeOf
}

// NewAndroid ...
func NewAndroid(andrType string) Android {
	return &android{
		typeOf: andrType,
	}
}
