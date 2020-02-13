package BFS

import "myrepo/Data_Structures/Binary_Tree/basic"

// 遍历树的级别顺序
func PrintLevelorder(root *basic.Node) [][]int {
	var res [][]int
	h := basic.Heigh(root)
	for i := 1; i <= h; i++ {
		var s []int
		PrintGivenLevel(root, i, &s)
		res = append(res, s)
	}
	return res
}

// 遍历给定级别
func PrintGivenLevel(root *basic.Node, level int, res *[]int) {
	if root == nil {
		return
	}
	if level == 1 {
		*res = append(*res, root.Data)
	} else if level > 1 {
		PrintGivenLevel(root.Left, level-1, res)
		PrintGivenLevel(root.Right, level-1, res)
	}
}
