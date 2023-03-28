package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet()
	set.Add("0")
	set.Add("1")
	set.Add("1")
	set.Add("1")
	set.Remove("0")
	list := set.List()
	fmt.Println(list)
}
