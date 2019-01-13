package main

import (
	"fmt"
	"learning/btree"
	"math/rand"
)

func main() {
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
