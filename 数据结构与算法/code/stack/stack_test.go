package stack

import "testing"

func TestArrayStack(t *testing.T) {
	stack := new(ArrayStack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push("000")
	stack.Pop()
	stack.Pop()
	stack.Push(5)
	stack.ShowAll()
}
func TestListStack(t *testing.T) {
	stack := NewListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	stack.ShowAll()
}
