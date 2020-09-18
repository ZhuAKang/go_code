package queue

import (
	"fmt"
	"testing"
)

func TestCircQueue(t *testing.T) {
	queue := InitCircQueue(10)
	queue.EnQueuEe(3)
	queue.EnQueuEe(5)
	queue.EnQueuEe(7)
	queue.EnQueuEe(9)
	queue.EnQueuEe(8)
	front, _ := queue.FrontQueue()
	rear, _ := queue.RearQueue()

	fmt.Println("此时队首元素是：", front)
	fmt.Println("此时队尾元素是：", rear)
	_, _ = queue.DeQueue()
	_, _ = queue.DeQueue()
	_, _ = queue.DeQueue()

	front, _ = queue.FrontQueue()
	rear, _ = queue.RearQueue()

	fmt.Println("此时队首元素是：", front)
	fmt.Println("此时队尾元素是：", rear)
	queue.Clear()

	front, ok1 := queue.FrontQueue()
	rear, ok2 := queue.RearQueue()

	fmt.Println("清空后队首元素是：", front, ok1)
	fmt.Println("清空后队尾元素是：", rear, ok2)
}
