package stack

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type ArrayStack struct {
	array []interface{}
	size  int
	lock  sync.Mutex
}

func (stack *ArrayStack) Push(data interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.array = append(stack.array, data)
	stack.size += 1
}
func (stack *ArrayStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		logrus.Warnln("empty stack, return nil")
		return nil
	}
	data := stack.array[stack.size-1]
	newArray := make([]interface{}, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray
	// 栈中元素数量-1
	stack.size = stack.size - 1
	return data
}

// 获取栈顶元素，但不出栈
func (stack *ArrayStack) Peek() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		logrus.Warnln("empty stack, return nil")
		return nil
	}
	data := stack.array[stack.size-1]
	return data
}
func (stack *ArrayStack) Size() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.size
}
func (stack *ArrayStack) ShowAll() {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	for _, v := range stack.array {
		fmt.Println(v)
	}
}
