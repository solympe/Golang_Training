package reverseList

import (
    "fmt"

    lf "github.com/solympe/Golang_Training/leetcode/reverseList/listNode"
)

func ReverseList(head *lf.ListFunctions) *lf.ListFunctions {
    next := lf.ListFunctions.GetNext(*head)

	if head == nil || next == nil {
		return head
	}

	revPoint := ReverseList(&next)
    fmt.Println("IN - 2")
	//head.Next.Next = head
	fmt.Println("AA", next.GetNext())

    lf.AfterNext(head)
    fmt.Println("O")
    //head.Next = nil
	lf.ListFunctions.NillNext(*head)
    fmt.Println("OUT")
	return revPoint
}
