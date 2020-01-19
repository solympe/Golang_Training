package main

import (
	"fmt"

	islands "github.com/solympe/Golang_Training/pkg/num-islands"
)

func main() {
	arr := islands.NewIslands([][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}})
	answer := arr.NumIslands()

	fmt.Println(answer)
}
