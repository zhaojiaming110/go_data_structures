// 在未排序的数组中searc、insert、delete
package Array

import "fmt"

// 在未排序的数组中、search操作可以通过从第一个元素到最后一个元素的线性遍历来执行。
func FindElement(arr []int, key int) int {
	for i, v := range arr {
		if v == key {
			return i
		}
	}
	return -1
}

// 在数组末尾插入
func InsertSorted(arr *[]int, key int) int {
	n := len(*arr)
	if n >= cap(*arr) {
		return n
	}
	*arr = append(*arr, key)
	return n
}

// Delete Operation
func DeleteElement(arr *[]int, key int) int {
	pos := FindElement(*arr, key)
	n := len(*arr)
	if pos == -1 {
		fmt.Println("Element not found")
		return n
	}
	(*arr) = append((*arr)[:pos], (*arr)[pos+1:]...)
	return n - 1
}
