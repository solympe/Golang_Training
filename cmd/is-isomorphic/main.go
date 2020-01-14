package main

import (
	"fmt"

	isom "github.com/solympe/Golang_Training/pkg/is-isomorphic"
)

func main() {
	word1 := isom.NewWorder("pam")
	word2 := isom.NewWorder("pum")

	fmt.Println("Answer: ", word1.IsIsomorphic(word2))
}
