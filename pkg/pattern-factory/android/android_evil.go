package android

type evilAndroid struct {
	typeOf string
}

// GetType ...
func (e *evilAndroid) GetType() string {
	return e.typeOf
}

// NewEvilAndroidManipulator ...
func NewEvilAndroidManipulator() AndroidManipulator {
	return &evilAndroid{
		typeOf: "evil",
	}
}
