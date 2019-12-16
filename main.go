package main

import (
	"fmt"

	rl "github.com/solympe/Golang_Training/leetcode/reverseList"
)

func main() {

	n4 := rl.ListNode {
		Val: 10,
		Next: nil,
	}

	n3 := rl.ListNode {
		Val: 10,
		Next: &n4,
	}

	n2 := rl.ListNode {
		Val: 10,
		Next: &n3,
	}

	n1 := rl.ListNode {
		Val: 10,
		Next: &n2,
	}


	fmt.Println(n1)


}
