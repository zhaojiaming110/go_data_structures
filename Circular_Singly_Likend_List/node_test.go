package node

import (
	"testing"
)

func Test_main(t *testing.T) {
	var last *Node
	// last->6
	addToEmpty(&last, 6)
	// last->6->1
	addBegin(&last, 1)
	// last->3 3->1->6
	addEnd(&last, 3)
	// last->3 3->1->9->6
	addAfter(last, 9, 1)
	traverse(last)
}
