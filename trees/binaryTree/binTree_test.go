// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/3 上午2:08

package binaryTree

import (
	"fmt"
	"github.com/zhaojiaming110/go_data_structures/Stacks/linkedlistStack"
	"testing"
)

func Test_binTree_Traverse(t *testing.T) {
	tree := New()
	tree.InsertAsRoot(2)
	tree.root.lChild = &binNode{
		data:1,
		lChild:&binNode{data:5},
		rChild:&binNode{data:4},
	}
	tree.root.rChild = &binNode{data:3}
	//         2
	//         |
	//     1	   3
	//   5	 4
	values := make([]interface{}, 0)
	tree.Traverse(TraLevel, func(data interface{}){
		values = append(values, data)
	})
	fmt.Println(values)
}

// 后序迭代遍历 实现特定visit(借助栈遍历出对应结果)
func Test_traPost_I(t *testing.T) {
	tree := New()
	tree.InsertAsRoot(2)
	tree.root.lChild = &binNode{
		data:1,
		lChild:&binNode{data:5},
		rChild:&binNode{data:4},
	}
	tree.root.rChild = &binNode{data:3}
	//         2
	//         |
	//     1	   3
	//   5	 4
	stack := linkedlistStack.New()
	tree.Traverse(TraPost_I, func(data interface{}){
		stack.Push(data)
	})
	fmt.Println(stack.Values())
}