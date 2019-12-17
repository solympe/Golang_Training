package main

import (
	"fmt"
	s "github.com/solympe/Golang_Training/leetcode/interselect/slice"
)

func main() {
	var first = s.NewSlice([]int{1, 2, 2, 1, 4, 6})
	var second = s.NewSlice([]int{2, 5, 4, 6})

	result := s.Intersect(first.GetSlice(), second.GetSlice())
	fmt.Println(result)
}
