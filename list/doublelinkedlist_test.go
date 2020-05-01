// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/1 下午4:08

package list

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	l := New(1,2,3,2,4,6,3,4)
	fmt.Println(l.Values())
	count := l.Deduplicate()
	fmt.Println(l.Values())
	fmt.Println(count)
}

func TestLists_Get(t *testing.T) {
	l := New(1,2,3)
	fmt.Println(l.Size())
	fmt.Println(l.Get(0))
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
	fmt.Println(l.Get(3))
}

func TestLists_Uniquify(t *testing.T) {
	l := New(1,1,2,2,2,3,4,4,5,5,6,6,6,7,8)
	fmt.Println(l.size)
	fmt.Println(l.Values())
	fmt.Println(l.Uniquify())
	fmt.Println(l.Values())
	fmt.Println(l.size)
}

func Test_selectionSort(t *testing.T) {
	l := New(1,2,2,1,4,2,7,6,3)
	l.Sort(SelectionSort, l.header.next, l.size)
	fmt.Println(l.Values())
}

func Test_insertionSort(t *testing.T) {
	l := New(1,2,2,1,4,2,7,6,3)
	insertionSort(l, l.header.next, l.size)
	fmt.Println(l.Values())
}

func Test_mergeSort(t *testing.T) {
	l := New(3,2,1,4,5,1,2,5,4)
	mergeSort(l, l.header.next, l.size)
	fmt.Println(l.Values())
}

func TestLists_Sort(t *testing.T) {
	l := New(3,2,1,4,5,1,2,5,4)
	l.Sort(InsetionSort, l.header.next, l.size)
	fmt.Println(l.Values())
	fmt.Println(l.Disordered())
}

func TestLists_Traverse(t *testing.T) {
	l := New(2,1,3,4)
	values := make([]interface{}, 0)
	l.Traverse(func(data interface{}) {
		values = append(values, data)
	})
	fmt.Println(values)
	fmt.Println(l.Disordered())
}