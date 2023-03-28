package tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	root := &BinaryTree{Data: "A"}
	root.Left = &BinaryTree{Data: "B"}
	root.Right = &BinaryTree{Data: "C"}
	root.Left.Left = &BinaryTree{Data: "D"}
	root.Left.Right = &BinaryTree{Data: "E"}
	root.Right.Left = &BinaryTree{Data: "F"}
	PostOrder(root)
	fmt.Println("----------------------------------")
	MidOrder(root)
	fmt.Println("----------------------------------")
	PreOrder(root)
	fmt.Println("----------------------------------")
	LayerOrder(root)
}
