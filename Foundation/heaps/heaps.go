package main

import (
	"fmt"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

// todo -> convert bst into min heap (love babbar)

func main() {
	input := []int64{4, 3, 2, 6}

	ropeCost := ropeMinCost(input)
	fmt.Println("rope cost::", ropeCost)

	input2 := []int{1, 2, 3, 4, 5}

	kthLargest := kthLargestSumSubarray(input2, 2)
	fmt.Println("kthLargest sum::", kthLargest)

	arrays1 := [][]int{
		{1, 3, 15, 17},
		{2, 4, 6, 8},
		{0, 9, 10, 11},
	}
	result1 := mergeKSortedArrays(arrays1, 3)
	fmt.Println("Test case 1 result:", result1)
}

func heapAndQueue() {
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

	// Show value in priority order by dequeuing
	fmt.Println("\nDequeueing from priority queue in order:")
	pqOrdered := priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(int) - a.(int)
	})

	for _, val := range []int{54, 53, 55, 52, 50} {
		pqOrdered.Enqueue(val)
	}

	for !pqOrdered.Empty() {
		val, _ := pqOrdered.Dequeue()
		fmt.Print(val, " ")
	}
	fmt.Println()
}
