package Stack

import "fmt"

type ArrayStack struct {
	Top      int
	Capacity int
	Data     []int
}

// create a stack.
func CreateStack(capacity int) *ArrayStack {
	return &ArrayStack{
		Top:      -1,
		Capacity: capacity,
		Data:     make([]int, capacity),
	}
}

// 当top等于最后一个索引时，Stack已满
func (s *ArrayStack) IsFull() bool {
	return s.Top == s.Capacity-1
}

// 当top等于-1时，Stack为空
func (s *ArrayStack) IsEmpty() bool {
	return s.Top == -1
}

// 入栈
func (s *ArrayStack) Push(item int) {
	if s.IsFull() {
		fmt.Println("Stack已满")
	}
	s.Top++
	s.Data[s.Top] = item
	fmt.Printf("%d pushed to stack\n", item)
}

// 出栈
func (s *ArrayStack) Pop() (item int) {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
	}
	item = s.Data[s.Top]
	s.Top--
	return
}

// 返回栈顶元素
func (s *ArrayStack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
	}
	return s.Data[s.Top]
}
