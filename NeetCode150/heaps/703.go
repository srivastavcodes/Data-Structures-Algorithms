package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type KthLargest struct {
	size    int
	minHeap *priorityqueue.Queue
}

func Constructor(k int, arr []int) KthLargest {
	pq := priorityqueue.NewWith(utils.IntComparator)
	for i := range arr {
		pq.Enqueue(arr[i])
	}
	for pq.Size() > k {
		pq.Dequeue()
	}
	return KthLargest{minHeap: pq, size: k}
}

func (kl *KthLargest) Add(val int) int {
	kl.minHeap.Enqueue(val)
	if kl.minHeap.Size() > kl.size {
		kl.minHeap.Dequeue()
	}
	curr, _ := kl.minHeap.Peek()
	return curr.(int)
}
