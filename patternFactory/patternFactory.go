package main

import "log"

type (
	Android interface {
		getType() string
		getName() string
		giveName(string)
	}

	evilAndroid struct {
		name   string
		typeOf string
	}

	kindAndroid struct {
		name   string
		typeOf string
	}
)

func (e evilAndroid) getName() string {
	return e.name
}

func (k kindAndroid) getName() string {
	return k.name
}

func (e evilAndroid) getType() string {
	return e.typeOf
}

func (k kindAndroid) getType() string {
	return k.typeOf
}

func (e *evilAndroid) giveName(name string) {
	e.name = name
}

func (k *kindAndroid) giveName(name string) {
	k.name = name
}

func AndroidFactory(typeOf string) Android {
	var product Android
	switch typeOf {
	case "evil":
		product = &evilAndroid{
			name:   "",
			typeOf: "evil",
		}
	case "kind":
		product = &kindAndroid{
			name:   "",
			typeOf: "kind",
		}
	default:
		log.Fatal("Error! Unknown type")
	}
	return product
}
