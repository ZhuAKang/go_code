package heap

import (
	"fmt"
	"testing"
)

func TestAHeap(t *testing.T) {
	heap := NewAHeap(11)
	heap.Insert(5)
	heap.Insert(11)
	heap.Insert(7)
	heap.Insert(2)
	heap.Insert(90)
	heap.Insert(45)
	heap.Insert(23)
	heap.Insert(8)
	heap.Insert(12)
	heap.Insert(28)
	heap.Insert(77)
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())
	fmt.Println(heap.Remove())

}
