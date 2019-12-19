package Stack

import (
	"fmt"
	"testing"
)

func Test_StackSlice(t *testing.T) {
	s := CreateSliceStack()
	fmt.Println(IsEmptySliceStack(s))
	PushSliceStack(s, 10)
	fmt.Println(IsEmptySliceStack(s))
	PushSliceStack(s, 20)
	PushSliceStack(s, 30)
	fmt.Println(s.Data)
	fmt.Println(PopSliceStack(s))
	fmt.Println(s.Data)
	fmt.Println(PeekSliceStack(s))
}
