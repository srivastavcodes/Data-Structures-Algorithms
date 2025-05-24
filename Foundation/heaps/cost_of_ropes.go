package main

import "github.com/emirpasic/gods/queues/priorityqueue"

func ropeMinCost(arr []int64) int64 {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		val1, val2 := a.(int64), b.(int64)
		switch {
		case val1 < val2:
			return -1
		case val1 > val2:
			return 1
		default:
			return 0
		}
	})
	for i := 0; i < len(arr); i++ {
		pq.Enqueue(arr[i])
	}
	var cost int64

	for pq.Size() > 1 {
		l, _ := pq.Dequeue()
		r, _ := pq.Dequeue()

		sum := l.(int64) + r.(int64)
		cost += sum

		pq.Enqueue(sum)
	}
	return cost
}
