// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午2:59

// Package Stacks provides an abstract stack interface.
package Stacks

// Interface Stack that all stacks implement
type Stack interface {
	Size() int
	Empty() bool
	Push(value interface{}) bool
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)
	Clear()
	Values() (values []interface{})
}
