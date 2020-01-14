package main

import (
	"fmt"

	d "github.com/solympe/Golang_Training/pkg/diff-ways-to-compute"
)

const stringIn = "0+3-4*4"

func main() {

	res := d.NewDataBox()
	result := res.DiffWaysToCompute(stringIn)

	fmt.Println("ANSWER: ", result)
}
