package reverseList

import lf "github.com/solympe/Golang_Training/leetcode/reverseList/listNode"

func ReverseList(head *lf.ListFunctions) *lf.ListFunctions {
    next := lf.ListFunctions.GetNext(*head)

	if head == nil || next == nil {
		return head
	}
	revPoint := ReverseList(&next)

	//head.Next.Next = head
	next.GetNext() = nil
	//nn := lf.ListFunctions.GetNext(*head)
    //nn = *head
    lf.AfterNext(head)

    //head.Next = nil
	lf.ListFunctions.NillNext(*head)

	return revPoint
}
