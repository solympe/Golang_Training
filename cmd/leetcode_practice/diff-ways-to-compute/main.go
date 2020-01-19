package main

import (
	"fmt"

	data "github.com/solympe/Golang_Training/pkg/leetcode_practice/diff-ways-to-compute"
)

const stringIn = "0+3-4*4"

func main() {
	dataBox := data.NewData(stringIn)

	fmt.Println(dataBox.Compute())
}
