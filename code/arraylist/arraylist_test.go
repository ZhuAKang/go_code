package arraylist

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	// 可以这样来
	list := NewArrayList()
	// 也可以这样定义接口对象
	// var list List = NewArrayList()
	list.Append(3)
	list.Append(7)
	list.Append(9)
	fmt.Println("源数组：", list.String())
	list.Insert(2, 10)
	fmt.Println("在数组的2号位置插入10，此时数组为", list.String())
	for i := 0; i < 10; i++ {
		list.Insert(2, 10)
	}
	fmt.Println(list.String())

	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		fmt.Print(item, " ")
	}
}
