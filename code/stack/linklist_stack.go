package stack

import (
	"code/linklist"
	"errors"
)

// LLStack 栈的链表实现
type LLStack struct {
	// 栈的数据域
	dataSource *linklist.LinkList
	// 容量
	cap int
}

// NewLLStack 新建一个使用链表实现的栈
func NewLLStack(cap int) *LLStack {
	stack := new(LLStack)
	stack.cap = cap
	stack.dataSource = linklist.NewLinkList()
	return stack
}

// Clear 清空栈
func (stack *LLStack) Clear() {
	// 重新分配内存就行了
	stack.dataSource = linklist.NewLinkList()
}

// Size 栈的大小
func (stack *LLStack) Size() int {
	return stack.dataSource.Size()
}

// Pop 栈顶元素出栈
func (stack *LLStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("栈为空，不存在栈顶元素")
	}
	// 获取链表尾最后一个元素（栈顶）
	value, _ := stack.dataSource.Get(stack.dataSource.Size())
	_ = stack.dataSource.Delete(stack.dataSource.Size())
	return value, nil
}

// Push 进栈
func (stack *LLStack) Push(value interface{}) error {
	if stack.IsFull() {
		return errors.New("栈已满，不能 push 元素进入")
	}
	// 未满直接追加数据
	stack.dataSource.Append(value)
	return nil
}

// IsEmpty 栈是否为空
func (stack *LLStack) IsEmpty() bool {
	if stack.Size() == 0 {
		return true
	}
	return false
}

// IsFull 栈是否已满
func (stack *LLStack) IsFull() bool {
	if stack.dataSource.Size() >= stack.cap {
		return true
	}
	return false
}
