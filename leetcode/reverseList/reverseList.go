package reverseList

import "fmt"

// ListNode is a definition for singly-linked list.
 type ListNode struct {
    Val int
    Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
    fmt.Println(head.Val)
    if head == nil || head.Next == nil {
        return head
    }
    revPoint := ReverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return revPoint
}