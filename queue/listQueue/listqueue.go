// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午8:32

package listQueue

type node struct {
	data	interface{}
	next 	*node
}

type queue struct{
	size	int
	front 	*node
	rear 	*node
}

func New() *queue {
	return &queue{
		size:  0,
		front: nil,
		rear:  nil,
	}
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) Empty() bool {
	return q == nil || q.front == nil
}

// Method Enqueue: add an value to the queue
//  It changes rear and size
// 	只允许从队列尾部插入
func (q *queue) Enqueue(value interface{}) bool {
	node := &node{
		data: value,
		next: nil,
	}
	if q.Empty() {
		q.front = node
		q.rear = node
		q.size++
		return true
	}
	// Add the new node at the end of queue and change rear
	q.rear.next = node
	q.rear = q.rear.next
	q.size++

	return true
}

// Method Dequeue:remove an value from the queue
//  It changes front and size
func (q *queue) Dequeue() (value interface{}, ok bool) {
	if q.Empty() {
		return nil, false
	}
	delNode := q.front
	q.front = q.front.next
	q.size--
	// 考虑出队后，队列成空的情况
	if q.front == nil {
		q.rear = nil
	}
	// 删除节点，使其不指向队列
	delNode.next = nil
	return delNode.data, true
}

func (q *queue) Front() (value interface{}, ok bool) {
	if q.Empty() {
		return nil, false
	}
	return q.front.data, true
}

func (q *queue) Values() (values []interface{}) {
	cur := q.front
	for cur != nil {
		values = append(values, cur.data)
		cur = cur.next
	}
	return values
}

func (q *queue) Clear() {
	q.size = 0
	q.front = nil
	q.rear = nil
}

