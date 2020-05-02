// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午3:31

package arrayStack

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	stack := New(6)
	fmt.Println(stack.Empty())
	stack.Push(1)
	stack.Push(3)
	fmt.Println(stack.Values())
	fmt.Println(stack.Empty())
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(2)
	fmt.Println(stack.Values())
	ok := stack.Push(2)
	fmt.Println(ok)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Values())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Values())
}