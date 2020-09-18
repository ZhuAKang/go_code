package queue

import (
	"errors"
)

// Queue 定义了一系列队列的函数接口
// 结构体只要实现了接口内的函数即可被当作队列来使用
// 可以是链式队列、基于数组的循环队列等等
type Queue interface {
	IsEmpty() bool                    // 队列是否为空
	IsFull() bool                     // 队列是否已满
	FrontQueue() (interface{}, error) // 队首元素
	RearQueue() (interface{}, error)  // 队尾元素
	EnQueuEe(interface{}) error       // 入队操作
	DeQueue() (interface{}, error)    // 出队操作
	Clear()                           // 清空队列元素
}

// CircQueue 循环队列
type CircQueue struct {
	dataSource []interface{}
	// 队首指向的就是队首的元素，而队尾指向的是最后一个元素的下一个位置
	// 这样在循环队列满的状态下，最多能存 maxNumber - 1 个元素
	front, rear int
	// 循环队列的容量
	maxNumber int
}

// InitCircQueue 初始化循环队列
// cap: 队列的容量
func InitCircQueue(cap int) *CircQueue {
	queue := new(CircQueue)
	queue.dataSource = make([]interface{}, cap, cap)
	queue.front, queue.rear = 0, 0
	queue.maxNumber = cap
	return queue
}

// IsEmpty 判断队列是否为空
func (queue *CircQueue) IsEmpty() bool {
	// 队首队尾的指针相同的时候说明队列已经空了
	return queue.front == queue.rear
}

// IsFull 判断队列是否为满
func (queue *CircQueue) IsFull() bool {
	if (queue.rear+1)%queue.maxNumber == queue.front {
		return true
	}
	return false
}

// FrontQueue 获取队首元素
func (queue *CircQueue) FrontQueue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("队列为空，不存在队首元素")
	}
	return queue.dataSource[queue.front], nil
}

// RearQueue 获取队尾元素
func (queue *CircQueue) RearQueue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("队列为空，不存在队尾元素")
	}
	return queue.dataSource[(queue.rear-1)%queue.maxNumber], nil
}

// EnQueuEe 入队操作
func (queue *CircQueue) EnQueuEe(value interface{}) error {
	if queue.IsFull() {
		return errors.New("队列已满，无法放入")
	}
	// 队列未满，放入元素
	queue.dataSource[queue.rear] = value
	queue.rear = (queue.rear + 1) % queue.maxNumber
	return nil
}

// DeQueue 出队操作
func (queue *CircQueue) DeQueue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("队列为空，无法出队")
	}
	// 队列非空，出队
	value := queue.dataSource[queue.front]
	queue.front = (queue.front + 1) % queue.maxNumber
	return value, nil
}

// Clear 清空队列元素
func (queue *CircQueue) Clear() {
	// 清空不用真的删除，只需要把队首队尾指针重置就行
	queue.front, queue.rear = 0, 0
}
