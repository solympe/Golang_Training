package main

import (
	"fmt"

	s "github.com/solympe/Golang_Training/pkg/pattern-singleton"
)

func main() {

	man1 := s.Singleton()
	man1.SetName("Vladimir")
	man1.SetSurname("Ivanov")
	man2 := s.Singleton()

	fmt.Println(man1)
	fmt.Println(man2)
}
