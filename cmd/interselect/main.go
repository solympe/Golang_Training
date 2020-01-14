package main

import (
	"fmt"

	s "github.com/solympe/Golang_Training/pkg/interselect"
)

func main() {
	var first = s.NewSlicer([]int{1, 2, 2, 1, 4, 6})
	var second = s.NewSlicer([]int{2, 5, 4, 6, 1})

	result := first.Intersect(first.GetSlice(), second.GetSlice())
	fmt.Println(result)
}
