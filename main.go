package main

import (
	"fmt"
	"learning/btree"
	list "learning/linked_list"
	"math/rand"
)

func printList() {
	fmt.Println("---- Linked list ----")

	list := new(list.List)
	list.Insert(5)
	list.Insert(4)

	fmt.Println(list.Display())
	fmt.Println(list.Reverse())
}

func printBtree() {
	fmt.Println("---- Btree ----")

	var tree = btree.NewTree(&btree.Node{Data: 5})

	for i := 0; i < 10; i++ {
		tree.Insert(btree.Int(rand.Intn(10 + 1)))
	}

	findMe := btree.Int(8)

	if tree.Find(findMe) {
		fmt.Printf("%d - is in the tree\n", findMe)
	} else {
		fmt.Println("there is no such value in the tree")
	}

	tree.Display()

}

func main() {

	printList()
	printBtree()

}
