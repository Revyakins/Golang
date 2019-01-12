package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Int int32

type Btree struct {
	root *Node
	sync.Mutex
}

type Node struct {
	Data  Int
	Left  *Node
	Right *Node
}

func NewTree(node *Node) *Btree {
	if node == nil {
		panic("bad node")
	}

	return &Btree{root: node}
}

func (tree *Btree) Insert(data Int) {

	tree.Lock()

	if data <= 0 {
		panic("value cannot be null")
	}

	node := &Node{Data: data}

	if tree.root == nil {
		tree.root = node
		return
	} else {
		insertNode(tree.root, node)
	}

	defer tree.Unlock()
}

func insertNode(node, newNode *Node) {
	if newNode.Data < node.Data {
		if node.Left == nil {
			node.Left = newNode
			return
		}
		insertNode(node.Left, newNode)
	}

	if newNode.Data > node.Data {
		if node.Right == nil {
			node.Right = newNode
			return
		}
		insertNode(node.Right, newNode)
	}
}

func (btree *Btree) Clear() {
	btree.root = nil
}

func (btree *Btree) Find(value Int) bool {

	btree.Lock()
	defer btree.Unlock()

	if btree.root == nil {
		panic("Btree is empty")
	}

	return search(btree.root, value)
}

func search(node *Node, value Int) bool {

	if node == nil {
		return false
	}

	switch {
	case node.Data > value:
		return search(node.Left, value)
	case node.Data < value:
		return search(node.Right, value)
	default:
		return true
	}
}

func stringify(node *Node, level int) {

	if node != nil {
		format := ""

		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++

		stringify(node.Left, level)

		fmt.Printf(format+"%d\n", node.Data)

		stringify(node.Right, level)
	}
}

func (btree *Btree) Display() {
	btree.Lock()
	defer btree.Unlock()

	fmt.Println("---------------------------")
	stringify(btree.root, 0)
	fmt.Println("---------------------------")
}

func main() {
	var tree = NewTree(&Node{Data: 5})

	for i := 0; i < 10; i++ {
		tree.Insert(Int(rand.Intn(10 + 1)))
	}

	findMe := Int(8)

	if tree.Find(findMe) {
		fmt.Printf("%d - is in the tree\n", findMe)
	} else {
		fmt.Println("there is no such value in the tree")
	}

	tree.Display()

}
