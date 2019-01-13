package linked_list

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(val interface{}) {
	node := &Node{
		next:  L.head,
		value: val,
	}

	if L.head != nil {
		L.head.prev = node
	}
	L.head = node

	l := L.head
	for l.next != nil {
		l = l.next
	}

	L.tail = l
}

func (L *List) Display() string {
	node := L.head

	for node != nil {
		fmt.Printf("%+v ->", node.value)
		node = node.next
	}
	return ""
}

func (L *List) Reverse() string {
	curr := L.head
	var prev *Node

	L.tail = L.head

	for curr != nil {
		next := curr.next

		curr.next = prev

		prev = curr
		curr = next
	}

	L.head = prev

	return ""
}
