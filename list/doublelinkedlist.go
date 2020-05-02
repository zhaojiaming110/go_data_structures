// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/1 下午3:37

package list

import (
	"fmt"
)

type node struct {
	data	interface{}
	pred 	*node
	next 	*node
}

type Lists struct {
	size 	int
	header 	*node
	trailer *node
}

// Sort支持的方式
const (
	SelectionSort = 1
	InsetionSort = 2
	MergeSort = 3
)

// 构造函数
func New(values ...interface{}) *Lists {
	list := &Lists{header:&node{}, trailer:&node{}}
	list.header.next = list.trailer
	list.trailer.pred = list.header

	if len(values) > 0 {
		list.InsertAsLast(values...)
	}

	return list
}

func (l *Lists) Size() int {
	return l.size
}

func (l *Lists) Values() []interface{} {
	values := make([]interface{}, l.size, l.size)
	p := l.First()
	for i := 0; p != l.trailer; i++ {
		values[i] = p.data
		p = p.next
	}
	return values
}

func (l *Lists) First() *node {
	if l.size == 0 {
		return nil
	}
	return l.header.next
}

func (l *Lists) Last() *node {
	if l.size == 0 {
		return nil
	}
	return l.trailer.pred
}

// Method Get 通过秩访问列表元素
//  时间复杂度O(index)
func (l *Lists) Get(index int) (interface{}, bool) {
	if l.size == 0 || !l.withInRange(index) {
		return nil, false
	}
	p := l.First()	// 从首节点出发
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p.data, true
}

func (l *Lists) InsertAsFirst(value interface{}) {
	newnode := &node{data:value, pred:l.header, next:l.header.next}
	l.header.next.pred = newnode
	l.header.next = newnode
	l.size++
}

func (l *Lists) InsertAsLast(values ...interface{}) {
	for _, v := range values {
		newnode := &node{data:v, pred:l.trailer.pred, next:l.trailer}
		l.trailer.pred.next = newnode
		l.trailer.pred = newnode
		l.size++
	}
}

// Method InsertBefore 元素e当作节点p的前驱插入
//  调用前请确保p为有效节点
func (l *Lists) InsertBefore(p *node, e interface{}) *node {
	newnode := &node{data:e, pred:p.pred, next:p}
	p.pred.next = newnode
	p.pred = newnode
	l.size++
	return newnode
}

// Method InsertAfter 元素e当作节点p的后驱插入
//  调用前请确保p为有效节点
func (l *Lists) InsertAfter(p *node, e interface{}) *node {
	newnode := &node{data:e, pred:p, next:p.next}
	p.next.pred = newnode
	p.next = newnode
	l.size++
	return newnode
}

// Method Remove:删除合法位置p处节点，返回其data
//  时间复杂度O(1)
func (l *Lists) Remove(p *node) interface{} {
	p.pred.next, p.next.pred = p.next, p.pred
	l.size--
	deleteNode(p)
	return p.data
}

// Func deleteNode:释放p节点
func deleteNode(p *node) {
	p.next = nil
	p.pred = nil
	p.data = nil
}

// Method Clear:删除链表
func (l *Lists) Clear() {
	l.size = 0
	l.header = nil
	l.trailer = nil
}

// Method Disordered:判断所有节点是否已按非降序排列
func (l *Lists) Disordered() bool {
	if l.size < 2 {
		return true
	}
	for p := l.First(); p.next != l.trailer; p = p.next {
		if p.data.(int) > p.next.data.(int) {
			return false
		}
	}
	return true
}

// Func Sort:列表排序统一入口
func (l *Lists) Sort(m int, p *node, n int) {
	switch m {
	case SelectionSort:
		selectionSort(l, p, n)
	case InsetionSort:
		insertionSort(l, p, n)
	case MergeSort:
		mergeSort(l, p, n)
	default:
		fmt.Println("Sort Method not recognized")
	}
}

// Method Find 查找目标元素e
func (l *Lists) Find(e interface{}) *node {
	if l.size == 0 {
		return nil
	}
	p := l.header.next
	for p != l.trailer {
		if e == p.data {
			return p
		}
		p = p.next
	}
	return nil
}

// Func find 在无序列表内节点p(可能是trailer)的n个真前驱中，找到等于e的最后者
//  调用前请确保p为列表中的有效节点，且n的值不会越界
func find(e interface{}, n int, p *node) *node {
	for n > 0 {	// 对于p的最近的n个前驱，从右向左
		p = p.pred
		if e == p.data {
			return p
		}
		n--
	}
	return nil	//p越出左边界意味着区间不含e，查找失败
}

// Func copyNodes:复制列表中自位置p起的n项
//  p合法且至少有n-1个真后继节点
func copyNodes(p *node, n int) *Lists {
	list := New()	// 创建头、尾哨兵节点并初始化
	for n > 0 {
		list.InsertAsLast(p.data)
		p = p.next
		n--
	}
	return list
}

// Method Deduplicate:剔除无序列表中的重复节点
//  时间复杂度O(N^2)
func (l *Lists) Deduplicate() int {
	if l.size < 2 {
		return 0
	}
	oldSize := l.size
	p := l.First()	// p从首节点开始, 遍历至末节点
	for r := 0; p != l.trailer; r++ {
		if q := find(p.data, r, p); q != nil {
			l.Remove(q)
			r--
		}
		p = p.next
	}
	return oldSize - l.size
}

// Method Uniquify:在有序列表中成批剔除重复元素
//  时间复杂度O(n)，只需遍历一次链表
func (l *Lists) Uniquify() int {
	if l.size < 2 {
		return 0
	}
	oldSize := l.size
	p := l.First()
	for q := p.next; q != l.trailer; q = p.next {
		if p.data == q.data {
			l.Remove(q)
		} else {
			p = q
		}
	}
	return oldSize - l.size
}

// Func selectionSort:列表的选择排序算法；对起始于位置p的n个元素进行排序
func selectionSort(l *Lists, p *node, n int) {	//valid(p) && rank(p) + n <= size
	if n < 2 {
		return
	}
	// Step 1. 确定待排序区间（head, tail) 将列表分为前缀无序子序列、中间待排序区间子序列、后缀有序子序列
	head := p.pred
	tail := p
	for i := 0; i < n; i++ {
		tail = tail.next	// 待排序区间为 （head, tail) 区间长度为n
	}
	for n > 1 {
		// Step 2. 找出待排序区间n中最大的节点max
		max := selectMax(head.next, n)
		// Step 3. 将最大节点max插入后缀有序子序列之前
		l.InsertBefore(tail, max.data)
		// Step 4. 将待排序区间中的最大节点max删除
		l.Remove(max)
		// Step 5. 待排序区间长度减1 后缀有序子序列增1
		n--
		tail = tail.pred
	}
}

// Func selectMax:从起始于位置p的n个元素中，找出最大的节点
func selectMax(p *node, n int) *node {
	max := p
	for cur := p; n > 1; n-- {
		cur = cur.next
		if cur.data.(int) > max.data.(int){
			max = cur
		}
	}
	return max
}

// Func insertionSort:列表的插入排序算法；对起始于位置p的n个元素进行排序
func insertionSort(l *Lists, p *node, n int) { //valid(p) && rank(p) + n <= size
	if n < 2 {
		return
	}
	for r := 0; r < n; r++ {
		findNode := search(p.data, r, p)
		l.InsertAfter(findNode, p.data)
		p = p.next
		l.Remove(p.pred)
	}
}

// Func search:在有序列表内部节点p的n个真前驱中，找到不大于e的最后者。
func search(e interface{}, n int, p *node) *node {
	// assert: 0 <= n <= rank(p) < size
	for ; n >= 0; n-- {
		p = p.pred
		if p.data == nil {	// p为header节点
			break
		}
		if p.data.(int) <= e.(int) {
			break
		}
	}
	return p
}

// Func mergeSort:列表的归并排序算法；对起始于位置p的n个元素进行排序
// 	时间复杂度 nlog(n)
func mergeSort(l *Lists, p *node, n int) *node { //valid(p) && rank(p) + n <= size
	if n < 2 {
		return p
	}
	// Step 1. 以中点为界限
	m := n >> 1
	q := p
	// Step 2. 均分列表
	for i := 0; i < m; i++ {
		q = q.next
	}
	// Step 3. 对前后子列表分别排序
	p = mergeSort(l, p, m)
	q = mergeSort(l, q, n-m)
	// Step 4. 归并
	p = merge(l, p, m, l, q, n-m)
	return p
}

// Func merge:有序列表的归并
func merge(l1 *Lists, p *node, n int, l2 *Lists, q *node, m int) *node {
	// assert: L1.valid(p) && rank(p) + n <= l1.size && l1.sorted(p, n)
	// 	L2.valid(q) && rank(p) + m <= l2.size && l2.sorted(p, n)
	// 注意: 在归并排序之类的场合， 有可能l1==l2 && rank(p) + n = rank(q)
	pp := p.pred	// 借助前驱（可能是header)
	for m > 0 {
		if n > 0 && p.data.(int) <= q.data.(int) {
			p = p.next
			if q == p {
				break
			}
			n--
		} else {
			l1.InsertBefore(p, q.data)
			q = q.next
			l2.Remove(q.pred)
			m--
		}
	}
	p = pp.next	// 归并后的新起点
	return p
}

// Func Traverse:利用函数对象机制的遍历
func (l *Lists) Traverse(visit func(interface{})) {
	p := l.First()
	if p == nil {
		return
	}
	for ; p != l.trailer; p = p.next {
		visit(p.data)
	}
}

// Method withInRange:检查索引
func (l *Lists) withInRange(index int) bool {
	return index >= 0 && index < l.size
}

