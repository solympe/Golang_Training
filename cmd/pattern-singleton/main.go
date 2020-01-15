package main

import (
	"fmt"

	singleton "github.com/solympe/Golang_Training/pkg/pattern-singleton"
)

func main() {
	man1 := singleton.MakeSingleton()
	man1.SetName("Vladimir")
	man1.SetSurname("Ivanov")

	man2 := singleton.MakeSingleton()

	fmt.Println(man1)
	fmt.Println(man2)
}
