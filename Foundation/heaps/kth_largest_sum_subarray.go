package main

import "github.com/emirpasic/gods/queues/priorityqueue"

func kthLargestSumSubarray(arr []int, k int) int {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	for i := range arr {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if pq.Size() < k {
				pq.Enqueue(sum)
			} else if top, _ := pq.Peek(); sum > top.(int) {
				pq.Dequeue()
				pq.Enqueue(sum)
			}
		}
	}
	result, _ := pq.Peek()
	return result.(int)
}
