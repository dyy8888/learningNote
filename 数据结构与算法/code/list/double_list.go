package list

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

// 节点
type Delement struct {
	data interface{}
	pre  *Delement
	next *Delement
}

// 双向链表
type DList struct {
	len    int
	header *Delement
	tail   *Delement
	lock   sync.Mutex
}

// 初始化新的链表
func NewDList() *DList {
	return &DList{
		len:    0,
		header: nil,
		tail:   nil,
		lock:   sync.Mutex{},
	}
}

// 向尾部添加数据
func (d *DList) Append(v interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	temp := &Delement{data: v}
	if d.isEmpty() {
		d.header = temp
		d.tail = temp
	} else {
		d.tail.next = temp
		temp.pre = d.tail
		d.tail = temp
	}
	d.len++
}

// 向头部添加数据
func (d *DList) AddHead(v interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	temp := &Delement{data: v}
	if d.isEmpty() {
		d.header = temp
		d.tail = temp
	} else {
		d.header.pre = temp
		temp.next = d.header
		d.header = temp
	}
	d.len++
}

// 从任意位置插入数据
// 如果链表本身为空，则忽视index，直接添加数据
// 如果输入的index不正确，则不做插入操作同时打印错误
func (d *DList) Insert(index int, v interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	temp := &Delement{data: v}
	if d.isEmpty() {
		logrus.Warnln("empty DList, ignore the index input")
		d.header = temp
		d.tail = temp
		d.len++
		return
	}
	if index >= d.len || index < 0 {
		logrus.Errorln("invalid index")
		return
	}
	if index == 0 {
		d.header.pre = temp
		temp.next = d.header
		d.header = temp
		d.len++
		return
	}
	if index == d.len-1 {
		d.tail.next = temp
		temp.pre = d.tail
		d.tail = temp
		d.len++
		return
	}
	if index <= (d.len / 2) {
		node := d.header
		for count := 0; count < index; count++ {
			node = node.next
		}
		temp.pre = node.pre
		temp.next = node
		node.pre = temp
	} else {
		node := d.tail
		for count := d.len - 1; count > index; count-- {
			node = node.pre
		}
		temp.pre = node.pre
		temp.next = node
		node.pre = temp
	}
	d.len++
}

// 从头部出队列
func (d *DList) Lpop() interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.isEmpty() {
		return nil
	}
	data := d.header.data
	d.header = d.header.next
	if d.header != nil {
		d.header.pre = nil
	} else {
		d.tail = nil
	}
	d.len--
	return data
}

// 从尾部出列
func (d *DList) Rpop() interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.isEmpty() {
		return nil
	}
	data := d.tail.data
	d.tail = d.tail.pre
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.header = nil
	}
	d.len--
	return data
}

// 根据index查找数据,头节点算是0号
func (d *DList) GetData(index int) interface{} {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.len == 0 || index >= d.len || index < 0 {
		return nil
	}
	if index <= (d.len / 2) {
		node := d.header
		for count := 0; count < index; count++ {
			node = node.next
		}
		return node.data
	} else {
		node := d.tail
		for count := d.len - 1; count > index; count-- {
			node = node.pre
		}
		return node.data
	}
}

// 根据index删除节点,头节点算是0号
func (d *DList) Delete(index int) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.len == 0 || index >= d.len || index < 0 {
		return
	}
	if index == 0 {
		d.header = d.header.next
		if d.header != nil {
			d.header.pre = nil
		} else {
			d.tail = nil
		}
		d.len--
		return
	}
	if index == d.len-1 {
		d.tail = d.tail.pre
		if d.tail != nil {
			d.tail.next = nil
		} else {
			d.header = nil
		}
		d.len--
		return
	}
	if index <= (d.len / 2) {
		node := d.header
		for count := 0; count < index; count++ {
			node = node.next
		}
		before := node.pre
		next := node.next
		before.next = next
		next.pre = before
		node.next = nil
		node.pre = nil
		d.len--
		return
	} else {
		node := d.tail
		for count := d.len - 1; count > index; count-- {
			node = node.pre
		}
		before := node.pre
		next := node.next
		before.next = next
		next.pre = before
		node.next = nil
		node.pre = nil
		d.len--
		return
	}
}

// 打印所有元素
func (d *DList) PrintAll() {
	if d.header == nil {
		return
	}
	i := 1
	curr := d.header
	for curr != nil {
		fmt.Printf("%d => %+v\n", i, curr.data)
		i++
		curr = curr.next
	}
}

// 判断链表是否为空
func (d *DList) isEmpty() bool {
	return d.len == 0
}
