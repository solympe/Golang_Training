package factory

type evilAndroid struct {
	name   string
	typeOf string
}

func (e *evilAndroid) getName() string {
	return e.name
}

func (e *evilAndroid) getType() string {
	return e.typeOf
}

func (e *evilAndroid) giveName(name string) {
	e.name = name
}

func NewAndroid(name string, typeOf string) Android {
	return &evilAndroid{
		name:   name,
		typeOf: typeOf,
	}
}

