// 使用切片实现堆栈 本质是个Slice
package Stack

import "fmt"

type SliceStack struct {
	Data []int
}

func CreateSliceStack() *SliceStack {
	return &SliceStack{
		Data: make([]int, 0),
	}
}

func IsEmptySliceStack(s *SliceStack) bool {
	return len(s.Data) == 0
}

// 入栈
func PushSliceStack(s *SliceStack, item int) {
	s.Data = append(s.Data, item)
	fmt.Printf("%d Pushed\n", item)
}

// 出栈
func PopSliceStack(s *SliceStack) (item int) {
	if len(s.Data) == 0 {
		fmt.Println("Stack is empty")
	}
	item = s.Data[len(s.Data)-1]
	s.Data = s.Data[:(len(s.Data) - 1)]
	return item
}

// 返回栈顶元素
func PeekSliceStack(s *SliceStack) int {
	if IsEmptySliceStack(s) {
		fmt.Println("Stack is empty")
	}
	return s.Data[len(s.Data)-1]
}
