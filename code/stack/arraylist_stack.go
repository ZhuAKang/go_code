package stack

import (
	"code/arraylist"
	"errors"
)

// ALStack 使用前面的 arraylist 来实现栈
type ALStack struct {
	arraylist *arraylist.ArrayList
	// ArrayList 里面已经有 size 了，这里有 cap 就行
	// 栈容量
	cap int
}

// InitALStack 初始化一个使用 ArrayList 实现的栈
func InitALStack(cap int) *ALStack {
	stack := new(ALStack)
	stack.arraylist = arraylist.NewArrayList()
	stack.cap = cap
	return stack
}

// Clear 清空栈
func (stack *ALStack) Clear() {
	// 调用 arraylist 的 Clear 函数重新分配内存即可
	stack.arraylist.Clear()
}

// Size 栈的大小
func (stack *ALStack) Size() int {
	return stack.arraylist.Size()
}

// Pop 栈顶元素出栈
func (stack *ALStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("栈内为空，无法弹出元素")
	}
	// 这里的弹出使用的都是 ArrayList 的函数
	value, _ := stack.arraylist.Get(stack.arraylist.Size() - 1)
	_ = stack.arraylist.Delete(stack.arraylist.Size() - 1)
	return value, nil
}

// Push 进栈
func (stack *ALStack) Push(value interface{}) error {
	if stack.IsFull() {
		return errors.New("栈已满，无法继续压入元素")
	}
	// 入栈直接把元素 追加到 arraylist 后面就行了
	stack.arraylist.Append(value)
	return nil
}

// IsEmpty 栈是否为空
func (stack *ALStack) IsEmpty() bool {
	return stack.arraylist.Size() == 0
}

// IsFull 栈是否已满
func (stack *ALStack) IsFull() bool {
	return stack.arraylist.Size() >= stack.cap
}
