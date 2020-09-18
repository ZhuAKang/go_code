package stack

import (
	"fmt"
	"testing"
)

func TestAStack(t *testing.T) {
	var stack Stack = InitAStack(10)
	stack.Push(1)
	stack.Push(5)
	stack.Push(7)
	stack.Push(9)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
func TestALStack(t *testing.T) {
	var stack Stack = InitALStack(10)
	stack.Push(1)
	stack.Push(5)
	stack.Push(7)
	stack.Push(9)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
