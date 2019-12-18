package Circular_Linked_List

import "fmt"

// A Circular Singly Linked List
type Node struct {
	Data int
	Next *Node
}

// Insertion in an empty List
func addToEmpty(last **Node, new_data int) {
	if *last != nil {
		return
	}
	new_node := &Node{Data: new_data}
	new_node.Next = new_node
	*last = new_node
}

// Insertion at the beginning of the list
func addBegin(last **Node, new_date int) {
	if *last == nil {
		addToEmpty(last, new_date)
		return
	}
	new_node := &Node{new_date, (*last).Next}
	(*last).Next = new_node
}

// Insertion at the end of the list
func addEnd(last **Node, new_data int) {
	if *last == nil {
		addToEmpty(last, new_data)
		return
	}
	new_node := &Node{new_data, (*last).Next}
	(*last).Next = new_node
	*last = new_node
}

// Insertion in between the nodes
func addAfter(last *Node, new_data int, target int) {
	if last == nil {
		return
	}
	p := last.Next
	for {
		if p.Data == target {
			new_node := &Node{new_data, p.Next}
			p.Next = new_node
			return
		}
		p = p.Next
		if p == last.Next {
			return
		}
	}
}

// traverse from begin position
func traverse(last *Node) {
	if last == nil {
		fmt.Println("List is empty!")
	}
	p := last.Next
	for {
		fmt.Printf("%d ", p.Data)
		p = p.Next
		if p == last.Next {
			return
		}
	}
}
