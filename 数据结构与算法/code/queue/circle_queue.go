package queue

import (
	"errors"
	"fmt"
	"sync"
)

type CircleQueue struct {
	maxsize int
	array   []string
	head    int
	tail    int
	lock    sync.Mutex
}

func NewCircleQueue(size int) *CircleQueue {
	temp := new(CircleQueue)
	temp.maxsize = size
	temp.array = make([]string, temp.maxsize, size)
	return temp
}
func (circle *CircleQueue) Add(v string) error {
	circle.lock.Lock()
	defer circle.lock.Unlock()
	//判断是否已经满了
	if circle.isFull() {
		return errors.New("full queue")
	}
	circle.array[circle.tail] = v
	circle.tail = (circle.tail + 1) % circle.maxsize
	return nil
}
func (circle *CircleQueue) Remove() (string, error) {
	circle.lock.Lock()
	defer circle.lock.Unlock()
	if circle.isEmpty() {
		return "", errors.New("empty queue")
	}
	value := circle.array[circle.head]
	circle.head = (circle.head + 1) % circle.maxsize
	return value, nil
}
func (circle *CircleQueue) ListCircleQueue() {
	fmt.Println("环形队列情况如下")
	//取出当前队列有多少个元素
	size := circle.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设计一个辅助变量，指向head
	tempHead := circle.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%s\t", tempHead, circle.array[tempHead])
		tempHead = (tempHead + 1) % circle.maxsize
	}
	fmt.Println()
}
func (circle *CircleQueue) isFull() bool {
	return (circle.tail+1)%circle.maxsize == circle.head
}
func (circle *CircleQueue) isEmpty() bool {
	return circle.tail == circle.head
}
func (circle *CircleQueue) Size() int {
	return (circle.tail + circle.maxsize - circle.head) % circle.maxsize
}
