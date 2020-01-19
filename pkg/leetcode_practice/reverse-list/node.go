package list

type nodeValue = int

// Node ...
type Node interface {
	Reverse() (head Node)
	Next() Node
	set(next Node)
}

type node struct {
	value nodeValue
	next  Node
}

// Reverse return reversed list
func (n *node) Reverse() (head Node) {
	if n == nil || n.next == nil {
		return n
	}
	head = n.next.Reverse()
	n.next.set(n)
	n.next = nil
	return head
}

// Next ...
func (n *node) Next() Node {
	return n.next
}

func (n *node) set(next Node) {
	n.next = next
}

// NewNode returns new element of list
func NewNode(value nodeValue, next Node) Node {
	return &node{
		value: value,
		next:  next,
	}
}
