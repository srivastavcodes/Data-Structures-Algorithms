package main

import (
	"math"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func rangeInKSortedList(arr [][]int) []int {
	mini, maxi := math.MaxInt, math.MinInt

	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		val1, val2 := a.(*Matrix), b.(*Matrix)
		switch {
		case val1.value < val2.value:
			return -1
		case val1.value > val2.value:
			return 1
		default:
			return 0
		}
	})
	for i := 0; i < len(arr); i++ {
		maxi = max(maxi, arr[i][0])
		mini = min(mini, arr[i][0])
		pq.Enqueue(&Matrix{value: arr[i][0], row: i, col: 0})
	}
	start, end := mini, maxi
	for !pq.Empty() {
		currIF, _ := pq.Dequeue()
		curr := currIF.(*Matrix)

		mini = curr.value
		if maxi-mini < end-start {
			start, end = mini, maxi
		}
		if curr.col+1 < len(arr[curr.row]) {
			maxi = max(maxi, arr[curr.row][curr.col+1])

			pq.Enqueue(&Matrix{row: curr.row, col: curr.col + 1,
				value: arr[curr.row][curr.col+1],
			})
		} else {
			break
		}
	}
	return []int{start, end}
}
