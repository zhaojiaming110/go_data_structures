package Queue

import "fmt"

// A structure to represent a queue
type ArrayQueue struct {
	Front, Rear, Size int
	Capacity          int
	Data              []int
}

func CreateArrayQueue(queue *ArrayQueue, cap int) *ArrayQueue {
	return &ArrayQueue{
		Front:    0,
		Rear:     cap - 1,
		Size:     0,
		Capacity: cap,
		Data:     make([]int, cap),
	}
}

// Queue is full when size becomes equal to the capacity
func IsFull(queue *ArrayQueue) bool {
	return queue.Size == queue.Capacity
}

// Queue is empty when size is 0
func IsEmpty(queue *ArrayQueue) bool {
	return queue.Size == 0
}

// Function to add an item to the queue.
// It changes rear and size
func EnArrayQueue(queue *ArrayQueue, item int) {
	if IsFull(queue) {
		fmt.Println("The Queue has Full")
		return
	}
	queue.Rear = (queue.Rear + 1) % (queue.Capacity)
	queue.Data[queue.Rear] = item
	queue.Size++
	fmt.Printf("%d enqueued to queue\n", item)
}

// Function to remove an item from queue.
// It changes front and size
func DeArrayQueue(queue *ArrayQueue) int {
	if IsEmpty(queue) {
		fmt.Println("The Queue is Empty")
		return -1
	}
	item := queue.Data[queue.Front]
	queue.Front = (queue.Front + 1) % (queue.Capacity)
	queue.Size--
	return item
}

// Function to get front of queue
func FrontArray(queue *ArrayQueue) int {
	if IsEmpty(queue) {
		fmt.Println("The Queue is Empty")
		return -1
	}
	return queue.Data[queue.Front]
}

// Function to get rear of queue
func RearArray(queue *ArrayQueue) int {
	if IsEmpty(queue) {
		fmt.Println("The Queue is Empty")
		return -1
	}
	return queue.Data[queue.Rear]
}
