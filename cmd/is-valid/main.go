package main

import (
	"fmt"

	val "github.com/solympe/Golang_Training/pkg/is-valid"
)

func main() {
	str := val.NewDataBoxValidator("())")

	fmt.Println(str.IsValid())
}
