package Queue

import (
	"fmt"
	"testing"
)

func Test_arraymain(t *testing.T) {
	var queue *ArrayQueue
	queue = CreateArrayQueue(queue, 100)
	fmt.Println(IsEmpty(queue))
	EnArrayQueue(queue, 10)
	fmt.Println(IsEmpty(queue))
	EnArrayQueue(queue, 20)
	EnArrayQueue(queue, 30)
	fmt.Println(FrontArray(queue))
	fmt.Println(RearArray(queue))
	fmt.Println(DeArrayQueue(queue))
	fmt.Println(FrontArray(queue))
}
