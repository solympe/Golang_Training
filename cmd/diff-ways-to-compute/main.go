package main

import data "github.com/solympe/Golang_Training/pkg/diff-ways-to-compute"

const stringIn = "0+3-4*4"

func main() {
	dataBox := data.NewData()

	dataBox.Compute(stringIn)
	dataBox.GetAnswer()
}
