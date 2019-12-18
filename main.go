package main

import (
	"fmt"
	td "github.com/solympe/Golang_Training/leetcode/numIslands/tdArray"
)

func main() {
	arr := td.NewtdArray([][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}})
	answer := arr.NumIslands(arr.GetSlice())

	fmt.Println(answer)
}
