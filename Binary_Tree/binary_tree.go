// 二叉树的遍历 前序、中序、后序 通过递归来实现
package Binary_Tree

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

//        1  <- root
//      /   \
//     2	 3
//   /  \    / \
//  4    5  nil nil
// 中序遍历（左根右）
func PrintInorder(root *Node) []int {
	if root == nil {
		return nil
	}
	t1 := PrintInorder(root.Left)
	t2 := PrintInorder(root.Right)
	return append(append(t1, root.Data), t2...)
}

//        1  <- root
//      /   \
//     2	 3
//   /  \    / \
//  4    5  nil nil
// 前序遍历（根左右）[1,2,4,5,3]
//func PrintPreorder(root *Node) []int {
//	var ss []int
//	if root == nil {
//		return nil
//	}
//	ss = append(ss, root.Data)
//	t1 := PrintPreorder(root.Left)
//	t2 := PrintPreorder(root.Right)
//	return append(ss, append(t1, t2...)...)
//}
func PrintPreorder(root *Node) []int {
	var res []int
	Preorder(root, &res)
	return res
}
func Preorder(root *Node, s *[]int) {
	if root != nil {
		*s = append(*s, root.Data)
		if root.Left != nil {
			Preorder(root.Left, s)
		}
		if root.Right != nil {
			Preorder(root.Right, s)
		}
	}
}

//        1  <- root
//      /   \
//     2	 3
//   /  \    / \
//  4    5  nil nil
// 后序遍历（左右根）【4，5，2，3, 1】
//func PrintPostorder(node *Node) []int {
//	if node == nil {
//		return nil
//	}
//	t1 := PrintPostorder(node.Left)
//	t2 := PrintPostorder(node.Right)
//	return append(append(t1, t2...), node.Data)
//}

func PrintPostorder(node *Node) []int {
	var res []int
	Postorder(node, &res)
	return res
}

func Postorder(node *Node, s *[]int) {
	if node != nil {
		if node.Left != nil {
			Postorder(node.Left, s)
		}
		if node.Right != nil {
			Postorder(node.Right, s)
		}
		*s = append(*s, node.Data)
	}
}
