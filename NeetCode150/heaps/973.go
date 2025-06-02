package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

func kClosest(points [][]int, k int) [][]int {
	minHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		val1, val2 := a.([]int)[0], b.([]int)[0]
		return val1 - val2
	})
	for _, point := range points {
		x, y := point[0], point[1]
		dist := x*x + y*y
		minHeap.Enqueue([]int{dist, x, y})
	}
	result := make([][]int, 0)
	for i := 0; i < k; i++ {
		item, _ := minHeap.Dequeue()
		point := item.([]int)
		result = append(result, []int{point[1], point[2]})
	}
	return result
}
