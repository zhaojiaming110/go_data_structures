package Stack

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	s := CreateStack(100)
	fmt.Println(s.IsEmpty())
	s.Push(1)
	fmt.Println(s.IsEmpty())
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Pop())
}
