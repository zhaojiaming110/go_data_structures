// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午4:46

package example

import (
	"github.com/zhaojiaming110/go_data_structures/Stacks"
)

func Convert(s Stacks.Stack, n int, base int) {
	var digit = []string{"0", "1", "2", "3", "4", "5", "6", "7",
		"8", "9", "A", "B", "C", "D", "E", "F"}
	for n > 0 {
		s.Push(digit[n%base])
		n = n/base
	}
}