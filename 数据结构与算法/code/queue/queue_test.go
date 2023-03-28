package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue1 := NewArrayQueue(10)
	queue2 := NewArrayQueue(10)
	queue1.Add("dyy")
	queue2.Add("fsr")
	queue1.ShowQueue()
	queue2.ShowQueue()
	queue1.Add("dyy")
	queue2.Add("fsr")
	queue1.ShowQueue()
	queue2.ShowQueue()
	fmt.Println(queue1.Remove())
	fmt.Println(queue2.Remove())
	queue1.ShowQueue()
	queue2.ShowQueue()
}
func TestCircleQueue(t *testing.T) {
	circleQueue := NewCircleQueue(5)
	circleQueue.Add("1")
	circleQueue.Add("2")
	circleQueue.Add("3")
	circleQueue.Add("4")
	circleQueue.Add("5")
	circleQueue.Add("6")
	circleQueue.Remove()
	circleQueue.Add("5")
	circleQueue.ListCircleQueue()
}
