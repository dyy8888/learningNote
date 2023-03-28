package queue

import (
	"fmt"
	"sync"
)

// 数组队列，先进先出
type ArrayQueue struct {
	maxsize int
	array   []interface{} // 底层切片
	size    int           // 队列的元素数量
	lock    sync.Mutex    // 为了并发安全使用的锁
}

func NewArrayQueue(size int) *ArrayQueue {
	queue := new(ArrayQueue)
	queue.maxsize = size
	return queue
}
func (arrQueue *ArrayQueue) Add(v interface{}) {
	arrQueue.lock.Lock()
	defer arrQueue.lock.Unlock()
	arrQueue.array = append(arrQueue.array, v)
	arrQueue.size += 1
}
func (arrQueue *ArrayQueue) Remove() interface{} {
	arrQueue.lock.Lock()
	defer arrQueue.lock.Unlock()
	// 队中元素已空
	if arrQueue.size == 0 {
		panic("empty")
	}
	// 队列最前面元素
	v := arrQueue.array[0]
	/*    直接原位移动，但缩容后继的空间不会被释放
	      for i := 1; i < queue.size; i++ {
	          // 从第一位开始进行数据移动
	          queue.array[i-1] = queue.array[i]
	      }
	      // 原数组缩容
	      queue.array = queue.array[0 : queue.size-1]
	*/
	// 创建新的数组，移动次数过多
	newArray := make([]interface{}, arrQueue.size-1)
	for i := 1; i < arrQueue.size; i++ {
		// 从老数组的第一位开始进行数据移动
		newArray[i-1] = arrQueue.array[i]
	}
	arrQueue.array = newArray

	// 队中元素数量-1
	arrQueue.size = arrQueue.size - 1
	return v
}
func (arrQueue *ArrayQueue) ShowQueue() {
	arrQueue.lock.Lock()
	defer arrQueue.lock.Unlock()
	for _, v := range arrQueue.array {
		fmt.Println(v)
	}
}
func (arrQueue *ArrayQueue) GetLength() int {
	arrQueue.lock.Lock()
	defer arrQueue.lock.Unlock()
	return arrQueue.size
}
