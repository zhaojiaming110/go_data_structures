// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/3/30 下午12:23

package utils

import "sort"

func Sort(values []interface{}, comparator Comparator) {
	sort.Sort(sortDate{values,comparator})
}

// implement sort.Interface
type sortDate struct{
	values		[]interface{}
	comparator	Comparator
}

func (s sortDate) Len() int {
	return len(s.values)
}

func (s sortDate) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}

func (s sortDate) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
