package list

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Element struct {
	data interface{}
	next *Element
}
type List struct {
	length int
	root   *Element
}

/*
单链表的初始化
1、生成新结点作为头结点，用头指针指向头结点
2、头结点的指针域置空
*/
func InitList() *List {
	element := new(Element)
	list := new(List)
	list.root = element
	list.length = 0
	return list
}

/*
在指定位置插入元素：根节点为0号节点,根节点一直存在，但是不保存数据
单链表的插入=>将值为e的新结点插入到表的第i个结点的位置上，即插入到结点a(i-1)与a(i)之间
1、查找结点a(i-1)并由指针p指向该结点
2、生成一个新结点*s
3、将新结点*s的数据域置为e
4、将新结点*s的指针域指向结点a(i)
5、将结点*p的指针域指向新结点*s
*/
func (l *List) InsertElem(index int, v interface{}) error {
	if index <= 0 || index > l.length {
		return errors.New("out of the length")
	}
	pre := l.root
	ele := &Element{data: v}
	if index == 1 {
		l.AddElem(v)
		return nil
	}
	for i := 0; i < index-1; i++ {
		pre = pre.next
	}
	ele.next = pre.next
	pre.next = ele
	l.length++
	return nil
}

// 添加在根节点后一个
func (l *List) AddElem(v interface{}) {
	ele := &Element{data: v}
	if l.isEmpty() {
		l.root.next = ele
	} else {
		ele.next = l.root.next
		l.root.next = ele
	}
	l.length++
}

// 添加在尾部
func (l *List) AppendElem(v interface{}) {
	ele := &Element{data: v}
	pre := l.root
	if l.isEmpty() {
		pre.next = ele
		l.length++
		return
	}
	for {
		if pre.next != nil {
			pre = pre.next
		} else {
			break
		}
	}
	pre.next = ele
	l.length++
}

// 获取链表长度
func (l *List) GetSize() int {
	return l.length
}

// 删除指定位置的节点
func (l *List) DeleteIndex(index int) {
	if l.isEmpty() {
		logrus.Error("empty list, need not delete")
		return
	}
	if index <= 0 || index > l.length {
		logrus.Errorln("out of the length")
		return
	}
	pre := l.root
	if index == 1 {
		next := pre.next.next
		del := pre.next
		pre.next = next
		del.next = nil
	} else {
		for count := 0; count < index-1; count++ {
			pre = pre.next
		}
		del := pre.next
		next := pre.next.next
		pre.next = next
		del.next = nil
	}
	l.length--
}

// 删除指定值的节点,仅删除找到的第一个值的节点
func (l *List) DeleteValue(v interface{}) {
	if l.isEmpty() {
		logrus.Error("empty list, need not delete")
		return
	}
	pre := l.root.next
	if pre.data == v {
		l.root.next = pre.next
		pre.next = nil
		return
	}
	for {
		if pre.next != nil {
			if pre.next.data == v {
				temp := pre.next.next
				target := pre.next
				pre.next = temp
				target.next = nil
				break
			} else {
				pre = pre.next
			}
		} else {
			logrus.Errorln("the data is not existing")
			break
		}
	}
	l.length--

}

// 查询是否包含指定值
func (l *List) IsContain(v interface{}) bool {
	pre := l.root
	for {
		if pre.next != nil {
			if pre.next.data == v {
				return true
			} else {
				pre = pre.next
			}
		} else {
			return false
		}
	}
}

// 查询指定位置的值,若不存在，返回空值
func (l *List) GetElem(index int) interface{} {
	if l.isEmpty() {
		return nil
	}
	if index <= 0 || index > l.length {
		return nil
	}
	pre := l.root
	for count := 0; count < index; count++ {
		pre = pre.next
	}
	return pre.data
}

// 遍历链表
func (l *List) ShowList() {
	if l.isEmpty() {
		logrus.Infoln("empty list")
	}
	pre := l.root.next
	for {
		if pre.next != nil {
			fmt.Print(pre.data, "=>")
			pre = pre.next
		} else {
			fmt.Println(pre.data, "=>", "nil")
			return
		}
	}
}
func (l *List) isEmpty() bool {
	return l.root.next == nil
}
