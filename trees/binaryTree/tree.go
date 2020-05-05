// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/3 上午1:42

package binaryTree

type BinaryTree interface {
	InsertAsRoot(data interface{})
	InsertAsLC(x *binNode, data interface{})
	InsertAsRC(x *binNode, data interface{})
	Traverse(model int, visit func( interface{}))
}