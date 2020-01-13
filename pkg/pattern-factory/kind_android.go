package factory

type kindAndroid struct {
	name   string
	typeOf string
}

func (k *kindAndroid) getName() string {
	return k.name
}

func (k *kindAndroid) getType() string {
	return k.typeOf
}

func (k *kindAndroid) giveName(name string) {
	k.name = name
}

func NewAndroid(name string, typeOf string) Android {
	return &kindAndroid{
		name:   name,
		typeOf: typeOf,
	}
}
