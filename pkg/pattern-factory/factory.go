package factory

import "log"

type Android interface {
	getType() string
	getName() string
	giveName(string)
}


func Factory(typeOf string) Android {
	var product Android
	p := NewAndroid()
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


