package queue

import (
	"code/linklist"
	"errors"
)

// LLQueue 队列的链表实现
type LLQueue struct {
	// 存放的数据域
	dataSource *linklist.LinkList
	// 队列的容量
	cap int
}

// NewLLQueue 初始化一个队列的链表实现
func NewLLQueue(cap int) *LLQueue {
	queue := new(LLQueue)
	queue.cap = cap
	queue.dataSource = linklist.NewLinkList()
	return queue
}

// IsEmpty 队列是否为空
func (queue *LLQueue) IsEmpty() bool {
	if queue.dataSource.Size() == 0 {
		return true
	}
	return false
}

// IsFull 队列是否已满
func (queue *LLQueue) IsFull() bool {
	if queue.dataSource.Size() == queue.cap {
		return true
	}
	return false
}

// FrontQueue 队首元素
func (queue *LLQueue) FrontQueue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("队列为空，不存在队首元素")
	}
	value, _ := queue.dataSource.Get(1)
	return value, nil
}

// RearQueue 队尾元素
func (queue *LLQueue) RearQueue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("队列为空，不存在队尾元素")
	}
	value, _ := queue.dataSource.Get(queue.dataSource.Size())
	return value, nil
}

// EnQueuEe 入队操作
func (queue *LLQueue) EnQueuEe(newVal interface{}) error {
	if queue.IsFull() {
		return errors.New("队列已满，无法入队")
	}
	queue.dataSource.Append(newVal)
	return nil
}

// DeQueue 出队操作
func (queue *LLQueue) DeQueue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("队列为空，无法出队")
	}
	value, _ := queue.dataSource.Get(1)
	_ = queue.dataSource.Delete(1)
	return value, nil
}

// Clear 清空队列元素
func (queue *LLQueue) Clear() {
	queue.dataSource = linklist.NewLinkList()
}
