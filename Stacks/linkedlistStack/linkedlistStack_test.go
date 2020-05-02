// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午4:02

package linkedlistStack

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	stack := New()
	fmt.Println(stack.Empty())
	stack.Push(1)
	stack.Push(3)
	fmt.Println(stack.size)
	fmt.Println(stack.Values())
	fmt.Println(stack.Empty())
	fmt.Println(stack.Push(3))
	fmt.Println(stack.Push(4))
	fmt.Println(stack.Push(5))
	fmt.Println(stack.Push(2))
	fmt.Println(stack.Values())
	ok := stack.Push(2)
	fmt.Println(ok)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Values())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Values())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Values())
	fmt.Println(stack.size)
}