// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午10:35

package binaryTree

// binNode节点增设了成员变量height.利用这些变量固然可以加速j静态的查询或搜索
// 但为保持这些变量的时效性，在所属二叉树发生结构性调整之后，这些变量都需要动态更新。
// 因此，究竟是否值得引入此类成员变量，必须权衡利弊。
type binNode struct {
	data	interface{}
	lChild, rChild *binNode
}

// 将e作为当前节点的左孩子插入该节点
func (n *binNode) insertAsLC(e interface{}) *binNode {
	newNode := &binNode{
		data:   e,
		lChild: nil,
		rChild: nil,
	}
	n.lChild = newNode
	return newNode
}

// 将e当作当前节点的有孩子插入该节点
func (n *binNode) insertAsRC(e interface{}) *binNode {
	newNode := &binNode{
		data:   e,
		lChild: nil,
		rChild: nil,
	}
	n.rChild = newNode
	return newNode
}
