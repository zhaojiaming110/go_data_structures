package Queue

import (
	"fmt"
	"testing"
)

func Test_nodemain(t *testing.T) {
	q := CreateQueueNode()
	EnqueueNode(q, 10)
	EnqueueNode(q, 20)
	EnqueueNode(q, 30)
	fmt.Println(q.Front.Key)
	fmt.Println(q.Rear.Key)
	fmt.Println(DequeueNode(q).Key)
	fmt.Println(q.Front.Key)
	fmt.Println(q.Rear.Key)
}
