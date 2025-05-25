package main

import "github.com/emirpasic/gods/queues/priorityqueue"

type Matrix struct {
	row   int
	col   int
	value int
}

func mergeKSortedArrays(arrs [][]int, k int) []int {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		val1, val2 := a.(Matrix), b.(Matrix)
		switch {
		case val1.value < val2.value:
			return -1
		case val1.value > val2.value:
			return 1
		default:
			return 0
		}
	})
	for i := 0; i < k; i++ {
		curr := Matrix{value: arrs[i][0], row: i, col: 0}
		pq.Enqueue(curr)
	}
	result := make([]int, 0)
	for !pq.Empty() {
		curr, _ := pq.Dequeue()
		result = append(result, curr.(Matrix).value)

		row, col := curr.(Matrix).row, curr.(Matrix).col
		if col+1 < len(arrs[row]) {
			next := Matrix{value: arrs[row][col+1], row: row, col: col + 1}
			pq.Enqueue(next)
		}
	}
	return result
}
