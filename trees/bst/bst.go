// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/3 下午3:05

package bst

import (
	"github.com/zhaojiaming110/go_data_structures/utils"
)

type node struct {
	lChild *node
	rChild *node
	key		interface{}
	value 	interface{}
}

type bsTree struct {
	root 		*node
	hot 		*node
	comparator 	utils.Comparator
}

func NewIntComparator() *bsTree {
	return &bsTree{root:nil, comparator:utils.IntComparator}
}

func (b *bsTree) Search(key interface{}) *node {
	return searchIn(b.root, key, &b.hot, b.comparator)
}

func searchIn(v *node, key interface{}, hot **node, compare utils.Comparator) *node {
	if  v == nil {
		return v
	}
	res := compare(key, v.key)
	if res == 0 {
		return v
	}
	*hot = v
	if res < 0 {
		return searchIn(v.lChild, key, hot, compare)
	} else {
		return searchIn(v.rChild, key, hot, compare)
	}
}

func (b *bsTree) Insert(key interface{}) *node {
	x := b.Search(key)
	if x == nil {
		x = &node{key:key}
		if utils.IntComparator(key, b.hot.key) > 0 {
			b.hot.rChild = x
		} else {
			b.hot.lChild = x
		}
	}
	return x
}
