// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午8:51

package listQueue

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	q := New()
	fmt.Println(q.Empty())
	fmt.Println(q.Front())
	fmt.Println(q.Dequeue())
	q.Enqueue(3)
	fmt.Println(q.Empty())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Values())
	q.Enqueue(4)
	fmt.Println(q.Size())
	fmt.Println(q.Values())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Values())
	fmt.Println(q.Front())
	q.Enqueue(1)
	q.Enqueue(7)
	q.Enqueue(3)
	q.Enqueue(5)
	fmt.Println(q.Values())
	fmt.Println(q.Enqueue(7))
}