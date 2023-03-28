package list

import (
	"fmt"
	"testing"
)

func TestSingleList(t *testing.T) {
	list := InitList()
	list.AddElem(1)
	list.AddElem(2)
	list.AppendElem(3)
	list.InsertElem(1, 5)
	list.DeleteIndex(3)
	list.DeleteValue(2)
	fmt.Println(list.GetSize())
	list.ShowList()
}
func TestDoubleList(t *testing.T) {
	list := NewDList()
	list.AddHead(1)
	// list.AddHead(2)
	// list.Append(3)
	// list.Insert(0, 9)
	list.Delete(0)
	list.PrintAll()
}
