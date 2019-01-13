package btree

import (
	"testing"
)

var tree *Btree

func init() {
	tree = NewTree(&Node{Data: 5})
}

func TestBtree_Find(t *testing.T) {

	tree.Insert(Int(3))
	tree.Insert(Int(8))

	findMe := Int(8)
	result := tree.Find(findMe)

	if !result {
		t.Error(
			"For", findMe,
			"expected", true,
			"got", result,
		)
	}
}
