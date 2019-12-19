package Stack

import (
	"fmt"
	"testing"
)

func Test_StackLinkList(t *testing.T) {
	var root *StackLinkedList
	fmt.Println(IsEmpty_L(root))
	Push_L(&root, 10)
	fmt.Println(IsEmpty_L(root))
	fmt.Println(Peek_L(root))
	fmt.Println(Pop_L(&root))
	fmt.Println(Peek_L(root))
}
