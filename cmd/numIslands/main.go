package main

import (
	"fmt"

	i "github.com/solympe/Golang_Training/pkg/numIslands"
)

func main() {
	arr := i.NewIslandCounter([][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}})
	answer := arr.NumIslands(arr.GetSlice())

	fmt.Println(answer)
}
