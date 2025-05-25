package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

func lastStoneWeight(stones []int) int {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(int) - a.(int)
	})
	for _, stone := range stones {
		pq.Enqueue(stone)
	}
	for pq.Size() > 1 {
		l, _ := pq.Dequeue()
		r, _ := pq.Dequeue()
		pq.Enqueue(l.(int) - r.(int))
	}
	top, _ := pq.Peek()
	return top.(int)
}
