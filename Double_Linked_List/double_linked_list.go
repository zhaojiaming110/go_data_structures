// Package double_linked_list implements a doubly linked list.

package double_linked_list

import (
	"fmt"
	"reflect"
)

// A node of Double Linked List
type Node struct {
	Data int
	Prev *Node
	Next *Node
}

// 在双链表前插入一个新的节点
func (head *Node) push(new_date int) *Node {
	// 1.创建新节点并初始化节点
	new_node := &Node{new_date, nil, head}
	if *head == (Node{}) {
		new_node.Next = nil
	} else {
		// 2.使头结点的前指针指向新节点
		head.Prev = new_node
	}
	// 返回新的头指针
	return new_node
}

// 在给定节点之后添加新的节点
func (pre *Node) insertafter(new_date int) {
	// 1.检查给定节点是否为空节点
	if *pre == (Node{}) {
		fmt.Println("The given previous node is null")
		return
	}
	// 2.创建新节点并初始化
	new_node := &Node{new_date, pre, pre.Next}
	// 3.改变前置节点的next，后置节点的prev
	pre.Next = new_node
	if new_node.Next != nil {
		new_node.Next.Prev = new_node
	}
}

// 在双链表末尾添加一个新的节点
func appendEnd(head **Node, new_date int) {
	new_node := &Node{Data: new_date, Next: nil}
	temp := *head

	// 给定链表为空链表的情况
	if **head == (Node{}) {
		new_node.Prev = nil
		*head = new_node
		return
	}

	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = new_node
	new_node.Prev = temp
}

// 在给定节点之前添加一个新的节点
func insertBefore(head **Node, next_node *Node, new_date int) {
	if *next_node == (Node{}) {
		fmt.Println("The given next_node cannot be NULL")
		return
	}

	// 创建新节点
	new_node := &Node{new_date, next_node.Prev, next_node}
	// 更改next_node.Prev
	next_node.Prev = new_node
	if new_node.Prev == nil {
		*head = new_node
	} else {
		new_node.Prev.Next = new_node
	}
}

// Delete a node in Double Linked List
func DeleteNode(head **Node, del *Node) {
	if **head == (Node{}) || *del == (Node{}) {
		return
	}
	// delete head node
	if *head == del {
		*head = del.Next
		//if del.Next == nil 相当于 delete Double linked list
	}
	if del.Next != nil {
		del.Next.Prev = del.Prev
	}
	if del.Prev != nil {
		del.Prev.Next = del.Next
	}
	if *head == nil {
		fmt.Println("the Double Linked List is Distroy")
	}
}

// 反转双链表
func reverse(head **Node) {
	temp := *head
	for temp.Next != nil {
		temp.Next, temp.Prev = temp.Prev, temp.Next
		temp = temp.Prev
	}
	temp.Next, temp.Prev = temp.Prev, temp.Next
	// 更改头指针
	*head = temp
}

func PrintList(head interface{}) {
	v := reflect.ValueOf(head)
	if v.IsNil() {
		return
	}

	vv := v.Interface().(*Node)
	if *vv == (Node{}) {
		fmt.Println("this linked list is empty")
		return
	}
	for vv != nil {
		fmt.Println(vv.Data)
		vv = vv.Next
	}
}
