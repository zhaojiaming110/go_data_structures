// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午8:03

package arrayQueue

type queue struct{
	front, rear, size, capacity int
	data	[]interface{}
}

func New(cap int) *queue {
	return &queue{
		front:    0,
		rear:     cap - 1,
		size:     0,
		capacity: cap,
		data:     make([]interface{}, cap),
	}
}

func isfull(q *queue) bool {
	return q.size == q.capacity
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) Empty() bool {
	return q.size == 0
}

// Method Enqueue: add an value to the queue
//  It changes rear and size
// 	只允许从队列尾部插入
func (q *queue) Enqueue(value interface{}) bool {
	if isfull(q) {
		return false
	}
	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = value
	q.size++
	return true
}

// Method Dequeue:remove an value from the queue
//  It changes front and size
func (q *queue) Dequeue() (value interface{}, ok bool) {
	if q.Empty() {
		return nil, false
	}
	value = q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return value, true
}

// Method Front:return front of queue
func (q *queue) Front() (value interface{}, ok bool) {
	if q.Empty() {
		return nil, false
	}
	return q.data[q.front], true
}

func (q *queue) Values() (values []interface{}) {
	index := q.front
	for i := 0; i < q.size; i++ {
		values = append(values, q.data[index])
		index = (index + 1) % q.capacity
	}
	return values
}

func (q *queue) Clear() {
	q.data = []interface{}{}
	q.size = 0
	q.capacity = 0
}
