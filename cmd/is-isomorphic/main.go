package main

import (
	"fmt"

	isom "github.com/solympe/Golang_Training/pkg/is-isomorphic"
)

func main() {
	word1 := isom.NewWorder("cat")
	word2 := isom.NewWorder("dog")

	fmt.Println("Answer: ", isom.IsIsomorphic(word1.GetData(), word2.GetData()))
}
