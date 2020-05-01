// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/3/30 下午9:12

package doublelinkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := New()
	t.Log(list.Empty())
	list.Add(1,2,3,4)
	t.Log(list.Empty(), list.size)
	t.Log(list.Values())
	list2 := New(3,4,5)
	t.Log(list2.Empty(), list2.size)
	t.Log(list2.Values())
}

func TestList_Values(t *testing.T) {
	list := New(1, 2, 3)
	t.Log(list.Values())
}

func TestList_Get(t *testing.T) {
	list := New(1, 2, 3, 4)
	for i := 0; i < list.size; i++ {
		t.Log(list.Get(i))
	}
	t.Log(list.Get(5))
}

func TestList_Add(t *testing.T) {
	list := New(1)
	t.Log(list.Values())
	list.Add(7,3,1)
	t.Log(list.Values())
}

func TestList_Remove(t *testing.T) {
	list := New(1,5,3,8)
	t.Log(list.Values())
	t.Log(list.Remove(1))
	t.Log(list.Values())
	t.Log(list.Remove(0))
	t.Log(list.Values())
	list.Add(3,8,9,1)
	t.Log(list.Values())
	t.Log(list.Remove(5))
	t.Log(list.Values())
	t.Log(list.Remove(5))
	t.Log(list.Values())
}

func TestList_Insert(t *testing.T) {
	//list := New(8,5,9,3)
	//t.Log(list.Values())
	//list.Insert(0,3,4)
	//t.Log(list.Values())
	//
	//list2 := New(2,5,1)
	//list2.Insert(1,3,4,5)
	//t.Log(list2.Values())
	//
	//list3 := New(8,7,9)
	//list3.Insert(2,1)
	//t.Log(list3.Values())
	//
	//list4 := New(8,1,7)
	//list4.Insert(3,5,6,7)
	//t.Log(list4.Values())

	list5 := New()
	t.Log(list5.first)
	list5.Insert(0,2,3)
	t.Log(list5.first)
	t.Log(list5.last)
	t.Log(list5.Values())
	t.Log(list5)
	t.Log(list5.size)
}

func TestList_Contains(t *testing.T) {
	list := New(2,3,1,6)
	t.Log(list.Contains(2))
	t.Log(list.Contains(2,3))
	t.Log(list.Contains(2,1,6))
	t.Log(list.Contains(2,1,6,3))
	t.Log(list.Contains(7))
}