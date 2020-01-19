package main

import (
	"fmt"

	data "github.com/solympe/Golang_Training/pkg/leetcode_practice/is-valid"
)

func main() {
	str := data.NewData("())")

	fmt.Println(str.IsValid())
}
