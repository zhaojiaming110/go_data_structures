// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午7:57

package queue

type Queue interface {
	Size() int
	Empty() bool
	Enqueue(value interface{}) bool
	Dequeue() (value interface{}, ok bool)
	Front()	(value interface{}, ok bool)
	Values() (values []interface{})
	Clear()
}
