// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午3:11

package arrayStack

type arrayStack struct{
	top			int
	capacity	int
	data 		[]interface{}
}

func New(cap int) *arrayStack {
	return &arrayStack{
		top:      -1,
		capacity: cap,
		data:     make([]interface{}, cap),
	}
}

func isFull(s *arrayStack) bool {
	return s.capacity == s.top + 1
}

func (s *arrayStack) Size() int {
	return s.top + 1
}

func (s *arrayStack) Empty() bool {
	return s.top == -1
}

func (s *arrayStack) Push(value interface{}) bool {
	if isFull(s) {
		return false
	}
	s.top++
	s.data[s.top] = value
	return true
}

func (s *arrayStack) Pop() (value interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	value = s.data[s.top]
	s.top--
	return value, true
}

func (s *arrayStack) Peek() (value interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	return s.data[s.top], true
}

func (s *arrayStack) Clear() {
	s.capacity = 0
	s.data = []interface{}{}
}

func (s *arrayStack) Values() (values []interface{}) {
	for i := 0; i <= s.top; i++ {
		values = append(values, s.data[i])
	}
	return values
}

