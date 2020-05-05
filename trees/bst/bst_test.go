// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/3 下午3:40

package bst

import (
	"fmt"
	"testing"
)

func TestNewIntComparator(t *testing.T) {
	bst := NewIntComparator()
	bst.root = &node{key:3}
	bst.root.rChild = &node{key:4}
	bst.root.lChild = &node{key:1}
	fmt.Println(bst.Insert(2))
	fmt.Println(bst.root.lChild.rChild)
}