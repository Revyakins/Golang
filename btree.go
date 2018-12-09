package main

import "log"

type Node struct {
	Data  int32
	Left  *Node
	Right *Node
}

func (node *Node) Insert(data int32) {

	if node == nil {
		node.Data = data
		return
	}

	if data < node.Data {
		if node.Left == nil {
			node.Left.Data = data
			return
		}
		node.Left.Insert(data)
	}

	if data > node.Data {
		if node.Right == nil {
			node.Right.Data = data
			return
		}
		node.Right.Insert(data)
	}

}

func (node *Node) Clear() {

}

func (node *Node) Find(value int32) {
	if node.Data < value {
		log.Fatalf("%s", "Values is large is treee value")
		return
	}

}
