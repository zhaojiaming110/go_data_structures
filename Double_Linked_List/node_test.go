package double_linked_list

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	head := &Node{}
	head = head.push(1)
	head = head.push(2)
	head = head.push(3)
	fmt.Println(head.Data)
}

func TestDeleteNode(t *testing.T) {
	head := &Node{}
	fmt.Println("---------------")
	PrintList(head)
	fmt.Println("---------------")
	head = head.push(1)
	head = head.push(2)
	head = head.push(3)
	head = head.push(4)
	PrintList(head)
	fmt.Println("---------------")
	DeleteNode(&head, head.Next)
	PrintList(head)
	fmt.Println("---------------")
	DeleteNode(&head, head)
	PrintList(head)
	fmt.Println("---------------")
	DeleteNode(&head, head.Next)
	PrintList(head)
	fmt.Println("---------------")
	DeleteNode(&head, head)
	PrintList(head)
}

func Test_reverse(t *testing.T) {
	head := &Node{}
	head = head.push(1)
	head = head.push(2)
	head = head.push(3)
	head = head.push(4)
	PrintList(head)
	fmt.Println("---------------")
	reverse(&head)
	PrintList(head)
}
