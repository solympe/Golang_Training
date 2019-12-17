package main

import (
	"fmt"
	s "github.com/solympe/Golang_Training/leetcode/interselect/slice"
	ss "github.com/solympe/Golang_Training/leetcode/interselect/SliceSolver"
)

func main() {

	var first = s.NewSlice([]int{1,2,2,1})
	var second = s.NewSlice([]int{2})



	result := s.Intersect(ss.SliceSolver.GetSlice(first), ss.SliceSolver.GetSlice(second))
	fmt.Println(result)
}
