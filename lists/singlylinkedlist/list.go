// Package singlylinkedlist implements the singly-linked list.
package singlylinkedlist

import (
	"github.com/zhaojiaming110/go_data_structures/utils"
)

// Struct of node
type node struct {
	value interface{}
	next  *node
}

// Step 1. Method-like COMMDS find
type List struct {
	first *node		// 首节点
	last  *node		// 末节点
	header 	*node	// 头哨兵节点
	trailer	*node	// 尾哨兵节点
	size  int
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
	list.header = nil
	list.trailer = nil
}

// 构造方法
func New(values ...interface{}) *List {
	list := &List{header:new(node), trailer:new(node)}
	list.header.next = list.trailer
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for i, node := 0, list.header.next; i < list.size; i, node = i+1, node.next {
		values[i] = node.value
	}
	return values
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	node := list.first
	for i := 0; i != index; i, node = i+1, node.next {}

	return node.value, true
}

func (list *List) Add(values ...interface{})  {
	for _, value := range values {
		newNode := &node{value:value}
		if list.size == 0 {
			list.first = newNode
			list.last = newNode
			list.header.next = list.first
			list.last.next = list.trailer
		} else {
			newNode.next = list.trailer
			list.last.next = newNode
			list.last = newNode
		}
		list.size++
	}
}

func (list *List) Remove(index int) (interface{}, bool){
	if !list.withinRange(index) {
		return nil, false
	}

	node :=	list.header.next		// 删除的节点
	beforeNode := list.header	// 删除节点的前一个节点
	for i := 0; i != index; i, node = i+1, node.next {
		beforeNode = node
	}

	// 此处关注下若删除的是首节点或末节点，要改变list.first和list.last的值，此时不会改变哨兵节点的值
	beforeNode.next = node.next
	value := node.value
	if index == list.size - 1 {
		list.last = beforeNode
	}
	list.first = list.header.next
	node = nil
	list.size--
	return value, true
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	newList := New(values...)	// 创建一个新的list
	list.size = list.size + len(values)
	beforeNode := list.header
	node := list.header.next
	for i := 0; i != index; i, node = i+1, node.next {
		beforeNode = node
	}
	beforeNode.next = newList.first
	newList.last.next = node
	list.first = list.header.next
}

func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}

	for _, value := range values {
		found := false
		for node := list.first; node != nil; node = node.next {
			if node.value == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) && i != j {
		var node1, node2 *node
		for n, currentNode := 0, list.first; node1 == nil || node2 == nil; n, currentNode = n + 1, currentNode.next {
			switch n {
			case i:
				node1 = currentNode
			case j:
				node2 = currentNode
			}
		}
		node1.value, node2.value = node2.value, node1.value
	}
}

func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		if list.size == index {


			list.Add(value)
		}
		return
	}
	foundNode := list.first
	for i := 0; i != index; {
		i++
		foundNode = foundNode.next
	}
	foundNode.value = value
}

// quickSort
func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	values := list.Values()
	utils.Sort(values, comparator)
	list.clear()
	list.Add(values...)
}

// 判断所有节点是否已按非降序排列
func (list *List) Disordered(comparator utils.Comparator) bool {
	if list.size < 2 {
		return true
	}

	value := list.first.value
	for i, cur := 0, list.first; i < list.size; i, cur = i+1, cur.next {
		if comparator(value, cur.value) > 0 {
			return false
		}
		value = cur.value
	}
	return true
}

// check index
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List) clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}
