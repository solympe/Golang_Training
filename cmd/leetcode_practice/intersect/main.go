package main

import (
	"fmt"

	slice "github.com/solympe/Golang_Training/pkg/leetcode_practice/intersect"
)

func main() {
	var firstSlice = slice.NewSlice([]int{1, 2, 2, 1, 4, 6})
	var secondSlice = slice.NewSlice([]int{2, 5, 4, 6, 1})

	result := firstSlice.Intersect(firstSlice.GetSlice(), secondSlice.GetSlice())
	fmt.Println(result)
}
