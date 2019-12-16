package Queue

// A linked list (LL) node to store a queue entry
type QNode struct {
	Key  int
	Next *QNode
}

// The queue, front stores the front node of LL and rear stores the
// last node of LL
type QueueNode struct {
	Front *QNode
	Rear  *QNode
}

// A utility function to create an empty queue
func CreateQueueNode() *QueueNode {
	return &QueueNode{}
}

// The function to add a key k to q
func EnqueueNode(q *QueueNode, key int) {
	temp_node := &QNode{key, nil}
	if q.Rear == nil {
		q.Front = temp_node
		q.Rear = temp_node
		return
	}
	// Add the new node at the end of queue and change rear
	q.Rear.Next = temp_node
	q.Rear = temp_node
}

// Function to remove a key from given queue q
func DequeueNode(q *QueueNode) *QNode {
	if q.Front == nil {
		return nil
	}
	// Store previous front and move front one node ahead
	temp_node := q.Front
	q.Front = q.Front.Next
	// If front becomes NULL, then change rear also as NULL
	if q.Front == nil {
		q.Rear = nil
	}
	return temp_node
}
