package Stack

import "fmt"

type StackLinkedList struct {
	Data int
	Next *StackLinkedList
}

func IsEmpty_L(s *StackLinkedList) bool {
	return s == nil
}

// 在堆栈中添加一个项目
func Push_L(root **StackLinkedList, data int) {
	new_node := &StackLinkedList{data, *root}
	*root = new_node
	fmt.Printf("%d pushed to stack\n", data)
}

// 从堆栈中删除一个元素
func Pop_L(root **StackLinkedList) (int, bool) {
	if IsEmpty_L(*root) {
		return -1, false
	}
	temp := (*root).Data
	*root = (*root).Next
	return temp, true
}

func Peek_L(root *StackLinkedList) (int, bool) {
	if IsEmpty_L(root) {
		return -1, false
	}
	return root.Data, true
}