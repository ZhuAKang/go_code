package tree

import (
	"testing"
)

func TestRBTree(t *testing.T) {
	tree := new(RBTree)
	node1 := RBNode{5, 5, nil, nil, 1, BLACK}
	tree.root = &node1
	tree.Insert(3, 3)
	tree.Insert(7, 7)
	tree.Insert(2, 2)
	tree.Insert(1, 1)
	tree.Insert(8, 8)
	// tree.DeleteMin()
	// tree.DeleteMax()
	tree.Delete(2)
}
