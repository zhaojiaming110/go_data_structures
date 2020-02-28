// Package double_linked_list implements a doubly linked list.

package List

// Node is an element of a linked list.
type Node struct {
	next, prev *Node
	list *List
	Value interface{}
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root Node
	len int
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.root.next = nil
	l.root.prev = nil
	l.len = 0

	return l
}

// New returns an initialized list.
func New() *List {return new(List).Init()}