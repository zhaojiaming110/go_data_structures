package basic

import (
	"fmt"
	"testing"
)

func TestHeigh(t *testing.T) {
	//        1  <- root
	//      /   \
	//     2	 3
	//   /  \    / \
	//  4    5  nil nil
	node := &Node{1, nil, nil}
	node.Left = &Node{2, nil, nil}
	node.Right = &Node{3, nil, nil}
	node.Left.Left = &Node{4, nil, nil}
	node.Left.Right = &Node{5, nil, nil}
	fmt.Println(Heigh(node))
	fmt.Println(MaxDepth(node))
}
