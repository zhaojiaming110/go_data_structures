// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午3:46

package linkedlistStack

type node struct {
	data 	interface{}
	next	*node
}

type stack struct{
	top		*node
	size 	int
}

func New() *stack {
	return &stack{
		top:  nil,
		size: 0,
	}
}

func (s *stack) Size() int {
	return s.size
}

func (s *stack) Empty() bool {
	return s.size == 0
}

func (s *stack) Push(value interface{}) bool {
	s.top = &node{
		data: value,
		next: s.top,
	}
	s.size++
	return true
}

func (s *stack) Pop() (value interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	value = s.top.data
	s.top = s.top.next
	s.size--
	return value, true
}

func (s *stack) Peek() (value interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	return s.top.data, true
}

func (s *stack) Clear() {
	s.size = 0
	s.top = nil
}

func (s *stack) Values() (values []interface{}) {
	for cur := s.top; cur != nil; cur = cur.next {
		values = append(values, cur.data)
	}
	return values
}
