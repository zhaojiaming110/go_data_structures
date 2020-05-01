package lists

import (
	"github.com/zhaojiaming110/go_data_structures/utils"
)

type node struct {
	data	interface{}
	pred	*node
	next	*node
}

//type List struct {
//	size 	int
//	first 	*node
//	last 	*node
//	header	*node
//	trailer *node
//}

type List interface {
	Size() int
	Empty() bool
	Clear()
	Values() []interface{}

	Get(index int) (interface{}, bool)
	Add(...interface{})
	Remove(index int) (interface{}, bool)
	Insert(index int, values ...interface{})
	Contains(values ...interface{}) bool
	Swap(i, j int)
	Set(index int, value interface{})
	Sort(comparator utils.Comparator)
	Disordered() bool
}