package main

import (
	"fmt"
	"sync"
)

type President struct {
	name    string
	surname string
}

var (
	once      sync.Once
	president *President
)

func GetPresident() *President {
	once.Do(func() {
		president = &President{}
	})
	return president
}

func main() {

	somethingMan1 := GetPresident()
	somethingMan1.name = "Alexandr"
	somethingMan1.surname = "Ivanov"

	somethingMan2 := GetPresident()

	fmt.Println(somethingMan1.name, somethingMan1,"=",somethingMan2.name, somethingMan2, "is", somethingMan1 == somethingMan2)
}
