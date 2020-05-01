// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/3/30 下午8:16

package doublelinkedlist

import (
	"github.com/zhaojiaming110/go_data_structures/utils"
)

type List struct {
	first 	*node
	last 	*node
	header 	*node
	trailer	*node
	size 	int
}

type node struct {
	value 	interface{}
	next 	*node
	prev	*node
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

func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for i, node := 0, list.first; i < list.size; i, node = i+1, node.next {
		values[i] = node.value
	}
	return values
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	node := list.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value, true
}

// 构造方法
func New(values ...interface{}) *List {
	list := &List{header:&node{prev:nil}, trailer:&node{next:nil}}
	list.header.next = list.trailer
	list.trailer.prev = list.header
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newNode := &node{value:value}
		if list.size == 0 {
			list.header.next = newNode
			list.trailer.prev = newNode
			newNode.prev = list.header
			newNode.next = list.trailer
			list.first = newNode
			list.last = newNode
		} else {
			list.last.next = newNode
			list.trailer.prev = newNode
			newNode.prev = list.last
			newNode.next = list.trailer
			// 更新list.last
			list.last = newNode
		}
		list.size++
	}
}

func (list *List) Remove(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	removeNode := list.first
	for i := 0; i < index; i++ {
		removeNode = removeNode.next
	}

	removeNode.prev.next = removeNode.next
	removeNode.next.prev = removeNode.prev
	list.size--

	// 重置首末节点
	if list.size == 0 {
		list.first = nil
		list.last = nil
	} else {
		list.first = list.header.next
		list.last = list.trailer.prev
	}

	return removeNode.value, true
}

// 在指定位置前插入
func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if list.size == index {
			list.Add(values...)
			return
		}
	}

	// 新建一个list
	insertList := New(values...)
	foundNode := list.first
	for i := 0; i < index; i++ {
		foundNode = foundNode.next
	}

	foundNode.prev.next = insertList.first
	insertList.first.prev = foundNode.prev
	insertList.last.next = foundNode
	foundNode.prev = insertList.last
	list.size = list.size + len(values)

	// 重置list.first
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
		for i, cur := 0, list.first; i < list.size; i, cur = i+1, cur.next {
			if value == cur.value {
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

func (List) Swap(i, j int) {
	panic("implement me")
}

func (List) Set(index int, value interface{}) {
	panic("implement me")
}

func (List) Sort(comparator utils.Comparator) {
	panic("implement me")
}

func (List) Disordered() bool {
	panic("implement me")
}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

