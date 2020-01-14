package android

type kindAndroid struct {
	typeOf string
}

// GetType ...
func (k *kindAndroid) GetType() string {
	return k.typeOf
}

// NewKindAndroidManipulator ...
func NewKindAndroidManipulator() AndroidManipulator {
	return &kindAndroid{
		typeOf: "kind",
	}
}
