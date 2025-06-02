package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

func findKthLargest(nums []int, k int) int {
	minHeap := priorityqueue.NewWith(utils.IntComparator)

	for _, num := range nums {
		minHeap.Enqueue(num)
		if minHeap.Size() > k {
			minHeap.Dequeue()
		}
	}
	top, _ := minHeap.Peek()
	return top.(int)
}
