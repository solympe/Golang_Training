package main

import (
	"fmt"

	i "github.com/solympe/Golang_Training/pkg/num-islands"
)

func main() {
	arr := i.NewIslandsCounter([][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}})
	answer := arr.NumIslands(arr.GetSlice())

	fmt.Println(answer)
}
