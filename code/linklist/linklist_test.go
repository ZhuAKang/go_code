package linklist

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {
	list := NewLinkList()
	list.Append(1)
	list.Append(3)
	list.Append(7)
	list.Append(9)
	list.Append(6)
	fmt.Println(list.String())
	list.Insert(2, 2)
	fmt.Println(list.String())
	fmt.Println(list.Get(2))
	list.Delete(2)
	fmt.Println(list.String())
	list.Set(5, 9)
	fmt.Println(list.String())
	list.Clear()
	fmt.Println(list.String())
}
