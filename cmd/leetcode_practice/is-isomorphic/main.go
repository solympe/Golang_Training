package main

import (
	"fmt"

	isom "github.com/solympe/Golang_Training/pkg/leetcode_practice/is-isomorphic"
)

func main() {
	word1 := isom.NewWord("pam")
	word2 := isom.NewWord("pum")

	fmt.Println("Answer: ", word1.IsIsomorphic(word2))
}
