package main

import (
	"fmt"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func main() {
	heap := Heap{vals: make([]int, 0, 10)}
	heap.insert(54)
	heap.insert(53)
	heap.insert(55)
	heap.insert(52)
	heap.insert(50)

	heap.printHeap()

	heap.deleteNode()

	heap.printHeap()

	fmt.Println()

	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(int) - a.(int)
	})

	pq.Enqueue(54)
	pq.Enqueue(53)
	pq.Enqueue(55)
	pq.Enqueue(52)
	pq.Enqueue(50)

	fmt.Println(pq.Values())

	pq.Dequeue()

	fmt.Println(pq.Values())
}
