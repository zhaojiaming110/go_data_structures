package Singly_Linked_List

import "fmt"

type Node struct {
	date interface{}
	next *Node
}

// func main() {
//     // Start with the empty list
//     head := &Node{}

//     // Insert 6. So linked list becomes 6->NULL
//     head = head.appendEnd(6)

//     // Insert 7 at the beginning.
//     // So linked list becomes 7->6->NULL
//     head = head.push(7)

//     // Insert 1 at the beginning.
//     // So linked list becomes 1->7->6->NULL
//     head = head.push(1)

//     // Insert 4 at the end. So
//     // linked list becomes 1->7->6->4->NULL
//     head.appendEnd(4)

//     // Insert 8, after 7. So linked
//     // list becomes 1->7->8->6->4->NULL
//     head.next.insertAfter(8)

//     fmt.Println("Created Linked list is: ")
//     head.printList()

//     fmt.Println(head.getCount())

//     return
// }

// This function prints contents of
// linked list starting from head
func (n *Node) PrintList() {
	if *n == (Node{}) {
		fmt.Println("The list is empty")
	}

	for n.next != nil {
		fmt.Println(n.date)
		n = n.next
	}
	fmt.Println(n.date)
}

// 给一个指向单链表头部的指针，在单链表的最前插入一个新的node
func (oldHead *Node) Push(new_date interface{}) (newHead *Node) {
	if *oldHead == (Node{}) {
		newHead = &Node{new_date, nil}
		return
	}
	newHead = &Node{new_date, oldHead}
	return
}

// 给定一个节点prev_node，在给定的节点之后添加新的节点
func (prev_node *Node) InsertAfter(new_date interface{}) {
	if *prev_node == (Node{}) {
		fmt.Println("the given previous node cannot be NULL")
		return
	}
	prev_node.next = &Node{new_date, prev_node.next}
}

// 给一个指向单链表头部的指针，在末端追加一个新的节点
func (head *Node) AppendEnd(new_date interface{}) (newHead *Node) {
	end := &Node{new_date, nil}
	last := head

	if *head == (Node{}) {
		newHead = end
		return newHead
	}

	for last.next != nil {
		last = last.next
	}

	last.next = end

	newHead = head

	return newHead
}

// 递归求单链表的长度
func (head *Node) GetCount2() int {
	if *head == (Node{}) {
		return 0
	}

	if head.next == nil {
		return 1
	}

	return 1 + head.next.GetCount2()
}

// func GetCount() func count(int) int {
// 	return func
// }

// 给定指向单链表的头指针和key,删除链表中key的第一个匹配项
func (head_ref *Node) DeleteNode(key interface{}) (head *Node) {
	temp := head_ref
	prev := &Node{}
	head = head_ref

	if *head_ref == (Node{}) {
		fmt.Println("空链表")
		return
	}

	if head_ref.next != nil && head_ref.date == key {
		head = head_ref.next
		return
	}

	for temp.next != nil && temp.date != key {
		prev = temp
		temp = temp.next
	}

	if temp.next == nil {
		fmt.Println("no this key")
	}

	prev.next = temp.next

	return

}

// 给定指向单链表的头指针和位置,删除链表中指定位置所在的节点
func (head_ref *Node) DeletePosition(position int) (head *Node) {
	head = head_ref
	temp := head_ref

	if *head_ref == (Node{}) {
		fmt.Println("链表为空")
		return
	}

	if position == 0 {
		head = head.next
		return
	}

	// 循环position-1次，直到获得指向前一个节点的指针temp
	for i := 0; temp.next != nil && i <= position-1; i++ {
		temp = temp.next
	}

	if temp.next == nil {
		fmt.Println("position超长")
		return
	}

	temp.next = temp.next.next

	return
}

// 删除链表
func (head *Node) DeleteList() *Node {
	head = nil
	return head
}

// 迭代求单链表的长度
func (head *Node) GetCount1() int {
	count := 0
	if *head == (Node{}) {
		fmt.Println("Link list is empty")
		return count
	}
	for head.next != nil {
		count++
		head = head.next
	}

	return count + 1
}
