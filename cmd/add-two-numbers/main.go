package main

import (
	"fmt"

	ln "github.com/solympe/Golang_Training/pkg/add-two-numbers/list-node"
)

func main() {
	list6 := ln.NewlistNode(5, nil)
	list5 := ln.NewlistNode(5, list6)
	list4 := ln.NewlistNode(5, list5)

	list3 := ln.NewlistNode(5, nil)
	list2 := ln.NewlistNode(5, list3)
	list1 := ln.NewlistNode(5, list2)

	res := ln.AddTwoNumbers(list1, list4)
	fmt.Println(res)

}
