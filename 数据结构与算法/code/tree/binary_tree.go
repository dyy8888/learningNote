package tree

import (
	"code/queue"
	"fmt"
)

type BinaryTree struct {
	Data  interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

// 先序遍历：先访问根节点，再访问左子树，最后访问右子树
func PreOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	fmt.Println("current data is ", tree.Data)
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

// 中序遍历：先访问左子树，再访问根节点，最后访问右子树。
func MidOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	MidOrder(tree.Left)
	fmt.Println(tree.Data)
	MidOrder(tree.Right)
}

// 后序遍历：先访问左子树，再访问右子树，最后访问根节点。
func PostOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	MidOrder(tree.Left)
	MidOrder(tree.Right)
	fmt.Println(tree.Data)
}

// 层次遍历
func LayerOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	queue := queue.NewArrayQueue(100)
	queue.Add(tree)
	for queue.GetLength() > 0 {
		element := queue.Remove()
		fmt.Println(element.(*BinaryTree).Data)
		if element.(*BinaryTree).Left != nil {
			queue.Add(element.(*BinaryTree).Left)
		}
		if element.(*BinaryTree).Right != nil {
			queue.Add(element.(*BinaryTree).Right)
		}
	}
}
