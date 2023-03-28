package stack

import (
	"code/list"
	"sync"

	"github.com/sirupsen/logrus"
)

// 使用单向链表
type ListStack struct {
	list list.List
	lock sync.Mutex
}

func NewListStack() *ListStack {
	link := list.InitList()
	return &ListStack{
		list: *link,
		lock: sync.Mutex{},
	}
}

func (stack *ListStack) Push(data interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.AppendElem(data)
	// fmt.Println(stack.list.GetSize())
}
func (stack *ListStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// fmt.Println(stack.list.GetSize())
	res := stack.list.GetElem(stack.list.GetSize())
	if res == nil {
		logrus.Warnln("empty get")
	}
	stack.list.DeleteIndex(stack.list.GetSize())
	return res
}
func (stack *ListStack) ShowAll() {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.ShowList()
}
