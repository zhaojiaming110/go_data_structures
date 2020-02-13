package basic

// 树节点
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// 递归计算树的高度
func Heigh(node *Node) int {
	if node == nil {
		return 0
	} else {
		lHeight := Heigh(node.Left)
		rHeight := Heigh(node.Right)
		if lHeight > rHeight {
			return lHeight + 1
		} else {
			return rHeight + 1
		}
	}
}

// 迭代获取树的高度
func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	// 创建一个队列，并将根节点入队
	queue := []*Node{root}
	depth := 0
	for len(queue) > 0 {
		size := len(queue)
		// 将该层的节点出队
		for i := 0; i < size; i++ {
			// 获取该节点
			node := queue[0]
			// 该节点出队
			queue = queue[1:]
			// 下层节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
