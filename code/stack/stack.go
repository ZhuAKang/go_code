package stack

import (
	"errors"
)

// Stack 栈的接口
type Stack interface {
	Clear()                    // 清空栈
	Size() int                 // 栈的大小
	Pop() (interface{}, error) // 栈顶元素出栈
	Push(interface{}) error    // 进栈
	IsEmpty() bool             // 栈是否为空
	IsFull() bool              // 栈是否已满
}

// AStack 使用数组实现栈
type AStack struct {
	dataSource []interface{} // 存放元素的切片
	cap        int           // 栈的最大容量
	size       int           //栈当前使用的容量
}

// InitAStack 初始化一个栈
func InitAStack(cap int) *AStack {
	stack := new(AStack)
	stack.size = 0
	stack.cap = cap
	stack.dataSource = make([]interface{}, cap, cap)
	return stack
}

// Clear 清空栈
func (stack *AStack) Clear() {
	// 重新分配内存即可
	stack.dataSource = make([]interface{}, stack.cap, stack.cap)
	stack.size = 0
}

// Size 栈的大小
func (stack *AStack) Size() int {
	return stack.size
}

// Pop 栈顶元素出栈
func (stack *AStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("栈内为空，无法弹出元素")
	}
	stack.size--
	value := stack.dataSource[stack.size]
	// 这里的弹出并不是真的将数据移出了切片，而是通过 size 来控制，之后再进来元素覆盖就行
	return value, nil
}

// Push 进栈
func (stack *AStack) Push(value interface{}) error {
	if stack.IsFull() {
		return errors.New("栈已满，无法继续压入元素")
	}
	stack.dataSource[stack.size] = value
	stack.size++
	return nil
}

// IsEmpty 栈是否为空
func (stack *AStack) IsEmpty() bool {
	return stack.size == 0
}

// IsFull 栈是否已满
func (stack *AStack) IsFull() bool {
	return stack.cap <= stack.size
}
