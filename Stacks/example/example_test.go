// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/2 下午5:00

package example

import (
	"fmt"
	"github.com/zhaojiaming110/go_data_structures/Stacks/arrayStack"
	"testing"
)

func TestConvert(t *testing.T) {
	//s := linkedlistStack.New()
	s := arrayStack.New(100)
	Convert(s, 100, 2)
	for !s.Empty() {
		v, _ := s.Pop()
		fmt.Print(v)
	}
	fmt.Println()
}