package datastructures

import "fmt"

func QueueSample() {
	fmt.Print(`
	Queue:
Queue is a linear data structure that follows the FIFO (First In First Out) principle.
The elements are inserted at the rear and deleted from the front.
Let's observe the example below:`)
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("\nQueue after enqueuing 1, 2, 3:", q.items)

	dequeued := q.Dequeue()
	fmt.Println("Dequeued item:", dequeued)
	fmt.Println("Queue after dequeuing:", q.items)
	fmt.Println("The complexity of this data structure is O(1) for both enqueue and dequeue operations.")
}

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
