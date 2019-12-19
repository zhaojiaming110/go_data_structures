// 二叉树遍历 前序、中序、后序 通过堆栈来实现
package Binary_Tree

type Node_s struct {
	Data *Node
	Next *Node_s
}

func IsEmptyNode_s(root *Node_s) bool {
	return root == nil
}

func PushNode_s(root **Node_s, data *Node) {
	newNode_s := &Node_s{data, *root}
	*root = newNode_s
	//fmt.Printf("%d Push\n", data.Data)
}

func PopNode_s(root **Node_s) (*Node, bool) {
	if IsEmptyNode_s(*root) {
		return nil, false
	}
	temp := (*root).Data
	*root = (*root).Next
	return temp, true
}

func InorderStackTraversal(root *Node) []int {
	if root == nil {
		return nil
	}
	//stack := &Node_s{}
	var stack *Node_s
	treeNode := root
	res := make([]int, 0)
	for treeNode != nil || !IsEmptyNode_s(stack) {
		for treeNode != nil {
			PushNode_s(&stack, treeNode)
			treeNode = treeNode.Left
		}
		if !IsEmptyNode_s(stack) {
			node, _ := PopNode_s(&stack)
			res = append(res, node.Data)
			treeNode = node.Right
		}
	}
	return res
}

// 使用堆栈来前序遍历二叉树
// 首先先将根节点压入栈中，然后弹出，依次将弹出的节点的右左节点压入栈中，注意顺序是"右 左"，因为栈是先入后出的，所以先让右节点进栈
func PreorderTraversal(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := &Node_s{root, nil}
	res := make([]int, 0)
	for !IsEmptyNode_s(stack) {
		node, _ := PopNode_s(&stack)
		res = append(res, node.Data)
		if node.Right != nil {
			PushNode_s(&stack, node.Right)
		}
		if node.Left != nil {
			PushNode_s(&stack, node.Left)
		}
	}
	return res
}

func PostorderTraversal(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := &Node_s{root, nil}
	res := make([]int, 0)
	for !IsEmptyNode_s(stack) {
		node, _ := PopNode_s(&stack)
		res = append(res, node.Data)
		if node.Left != nil {
			PushNode_s(&stack, node.Left)
		}
		if node.Right != nil {
			PushNode_s(&stack, node.Right)
		}
	}
	// 反转数组
	res_len := len(res)
	result := make([]int, res_len)
	for _, v := range res {
		result[res_len-1] = v
		res_len--
	}
	return result
}
