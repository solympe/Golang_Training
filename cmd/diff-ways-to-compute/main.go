package main

import (
	"fmt"

	data "github.com/solympe/Golang_Training/pkg/diff-ways-to-compute"
)

const stringIn = "0+3-4*4"

func main() {
	dataBox := data.NewData()

	fmt.Println(dataBox.Compute(stringIn))
}
