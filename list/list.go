// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/1 下午3:18

// Package list provides an abstract List interface.
package list

// Interface of List that all lists implements.
type List interface {
	Size() int
	Values() []interface{}
	First() *node
	Last() *node
	Get(index int) (interface{}, bool)
	InsertAsFirst(interface{})
	InsertAsLast(values ...interface{})
	InsertBefore(p *node, e interface{}) *node
	InsertAfter(p *node, e interface{})	*node
	Remove(p *node) interface{}
	Clear()
	Disordered() bool
	Sort(m int, p *node, n int)
	Find(e interface{}) *node
	Deduplicate() int
	Uniquify() int
	Traverse(visit func(interface{}))
}
