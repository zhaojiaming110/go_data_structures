// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午11:07

package binaryTree

import (
	"fmt"
	"github.com/zhaojiaming110/go_data_structures/Stacks"
	"github.com/zhaojiaming110/go_data_structures/Stacks/linkedlistStack"
	"github.com/zhaojiaming110/go_data_structures/queue/listQueue"
)

type binTree struct {
	root 	*binNode
}

const (
	TraPre_R = 1	//先序递归
	TraIn_R = 2	//中序递归
	TraPost_R = 3	//后序递归
	TraPre_I1 = 4 	//先序迭代1
	TraPre_I2 = 5	//先序迭代2
	TraIn_I = 6		// 中序迭代
	TraPost_I = 7	// 后序迭代
	TraLevel = 8	// 层次遍历
)

func New() *binTree {
	return &binTree{}
}

// 将data当作根节点插入空的二叉树
func (bt *binTree) InsertAsRoot(data interface{}) {
	bt.root = &binNode{data:data}
}

// data插入为右的做孩子
func (bt *binTree) InsertAsLC(x *binNode, data interface{}) *binNode {
	x.insertAsLC(data)
	return x.lChild
}

// data插入为x的左孩子
func (bt *binTree) InsertAsRC(x *binNode, data interface{}) *binNode {
	x.insertAsRC(data)
	return x.lChild
}

// 遍历
func (bt *binTree) Traverse(model int, visit func(interface{})) {
	switch model {
	case TraPre_R:
		traPre_R(bt.root, visit)
	case TraIn_R:
		traIn_R(bt.root, visit)
	case TraPost_R:
		traPost_R(bt.root, visit)
	case TraPre_I1:
		traPre_I1(bt.root, visit)
	case TraPre_I2:
		traPre_I2(bt.root, visit)
	case TraIn_I:
		traIn_I(bt.root, visit)
	case TraPost_I:
		traPost_I(bt.root, visit)
	case TraLevel:
		traLevel(bt.root, visit)
	default:
		fmt.Println("Traverse Method not recognized")
	}
}

// 先序递归遍历
func traPre_R(node *binNode, visit func(interface{})) {
	if node == nil {
		return
	}
	visit(node.data)
	traPre_R(node.lChild, visit)
	traPre_R(node.rChild, visit)
}

// 中序递归遍历
func traIn_R(node *binNode, visit func(interface{})) {
	if node == nil {
		return
	}
	traIn_R(node.lChild, visit)
	visit(node.data)
	traIn_R(node.rChild, visit)
}

// 后序递归遍历
func traPost_R(node *binNode, visit func(interface{})) {
	if node == nil {
		return
	}
	traPost_R(node.lChild, visit)
	traPost_R(node.rChild, visit)
	visit(node.data)
}

// 先序遍历版本1 此版本不容易推广到中序和后序遍历两个版本
func traPre_I1(node *binNode, visit func(interface{})) {
	stack := linkedlistStack.New()
	stack.Push(node)
	for !stack.Empty() {
		cur, _ := stack.Pop()
		curnode := cur.(*binNode)
		visit(curnode.data)
		if curnode.rChild != nil {
			stack.Push(curnode.rChild)
		}
		if curnode.lChild != nil {
			stack.Push(curnode.lChild)
		}
	}
}

func traPre_I2(node *binNode, visit func(interface{})) {
	stack := linkedlistStack.New()
	for {
		visitAlongLeftBranch(node, visit, stack)
		if stack.Empty() {
			break
		}
		data, _ := stack.Pop()
		node = data.(*binNode)
	}
}

func visitAlongLeftBranch(node *binNode, visit func(interface{}), s Stacks.Stack) {
	for node != nil {
		visit(node.data)
		if node.rChild != nil {
			s.Push(node.rChild)
		}
		node = node.lChild
	}
}

func goAlongLeftBranch(node *binNode, s Stacks.Stack) {
	// 反复地入栈，沿左分支深入
	for node != nil {
		s.Push(node)
		node = node.lChild
	}
}

func traIn_I(node *binNode, visit func(interface{})) {
	stack := linkedlistStack.New()	// 辅助栈
	for {
		goAlongLeftBranch(node, stack)	// 从当前节点出发，逐批入栈
		if stack.Empty() {
			break	// 直至所有节点遍历完毕
		}
		data, _ := stack.Pop()	// 弹出栈顶元素进行遍历
		curNode := data.(*binNode)
		visit(curNode.data)
		if curNode.rChild != nil { // 若当前节点右子树不为空，将控制权转向为当前节点的右子树
			node = curNode.rChild
		}
	}
}

func traPost_I(node *binNode, visit func(interface{})) {
	stack := linkedlistStack.New()
	stack.Push(node)
	for !stack.Empty() {
		data, _ := stack.Pop()
		curNode := data.(*binNode)
		visit(curNode.data)
		if curNode.lChild != nil {
			stack.Push(curNode.lChild)
		}
		if curNode.rChild != nil {
			stack.Push(curNode.rChild)
		}
	}
}

// 二叉树的层次遍历
func traLevel(node *binNode, visit func(interface{})) {
	queue := listQueue.New()
	queue.Enqueue(node)	// 先将根节点入队
	for !queue.Empty() {
		cur, _ := queue.Dequeue()
		curNode := cur.(*binNode)
		visit(curNode.data)
		if curNode.lChild != nil {
			queue.Enqueue(curNode.lChild)
		}
		if curNode.rChild != nil {
			queue.Enqueue(curNode.rChild)
		}
	}
}



