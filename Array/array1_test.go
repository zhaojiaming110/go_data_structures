package Array

import (
	"fmt"
	"testing"
)

func Test_array1(t *testing.T) {
	arr := make([]int, 0, 100)
	for i := 0; i < 3; i++ {
		arr = append(arr, i+1)
	}
	fmt.Println(FindElement(arr, 2))
	fmt.Println(InsertSorted(&arr, 4))
	fmt.Println(arr)
	fmt.Println(DeleteElement(&arr, 4))
	fmt.Println(arr)
}
