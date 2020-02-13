package BFS

import (
	"fmt"
	"github.com/zhaojiaming110/go_data_structures/Binary_Tree/basic"
	"testing"
)

func TestPrintLevelorder(t *testing.T) {
	//        1  <- root
	//      /   \
	//     2	 3
	//   /  \    / \
	//  4    5  nil nil
	node := &basic.Node{1, nil, nil}
	node.Left = &basic.Node{2, nil, nil}
	node.Right = &basic.Node{3, nil, nil}
	node.Left.Left = &basic.Node{4, nil, nil}
	node.Left.Right = &basic.Node{5, nil, nil}
	fmt.Println(PrintLevelorder(node))
}
