package main

import (
	"fmt"
	is "github.com/solympe/Golang_Training/leetcode/interselect"

)

func main() {

	first := is.NewSlice([]int{1,2,2,1})
	second := is.NewSlice([]int{2})


	result := is.Intersect(first.GetSlice(), second.GetSlice())
	fmt.Println(result)
}
