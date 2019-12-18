package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := map[*ListNode]int{}
	for head != nil && head.Next != nil {
		_, ok := m[head.Next]
		if ok == true {
			return true
		}
		m[head.Next] = head.Val
		head = head.Next
	}
	return false
}

func main() {

	node2 := ListNode{
		Val:  10,
		Next: nil,
	}

	node1 := ListNode{
		Val:  10,
		Next: &node2,
	}

	fmt.Println(hasCycle(&node1))
}
