package main

import "github.com/emirpasic/gods/queues/priorityqueue"

func kthSmallest(arr []int, k int) int {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(int) - a.(int)
	})
	for i := 0; i < k; i++ {
		pq.Enqueue(arr[i])
	}
	for i := k; i < len(arr); i++ {
		if val, _ := pq.Peek(); arr[i] < val.(int) {
			pq.Dequeue()
			pq.Enqueue(arr[i])
		}
	}
	res, _ := pq.Peek()
	return res.(int)
}
