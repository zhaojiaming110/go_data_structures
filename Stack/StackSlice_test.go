package Stack

import (
	"fmt"
	"testing"
)

func Test_StackSlice(t *testing.T) {
	s := CreateSliceStack()
	fmt.Println(IsEmptySliceStack(s))
	PushSliceStack(s, 10)
	fmt.Println(IsEmptySliceStack(s))
	PushSliceStack(s, 20)
	PushSliceStack(s, 30)
	fmt.Println(s.Data)
	fmt.Println(PopSliceStack(s))
	fmt.Println(s.Data)
	fmt.Println(PeekSliceStack(s))
}

func Test_Slice(t *testing.T) {
	var stack []int
	// IsEmpty
	if len(stack) == 0 {
		fmt.Println("The Stack is Empty")
	}
	// push操作
	stack = append(stack, 10)
	stack = append(stack, 20)
	// Peek
	item := stack[len(stack)-1]
	fmt.Println(item)
	// 出栈
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}
