package main

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

func (L *List) Display() {
	node := L.head

	for node != nil {
		fmt.Printf("%+v ->", node.value)
		node = node.next
	}
	fmt.Println()
}

func (L *List) Reverse() {
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
	L.Display()
}

func main() {
	var list = new(List)
	list.Insert(5)
	list.Insert(4)

	list.Display()
	list.Reverse()
}
