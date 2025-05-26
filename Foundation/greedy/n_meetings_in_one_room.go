package main

import (
	"sort"
)

// activity selection problem

func maxMeetings(start, end []int, length int) int {
	pair := make([][]int, 0)
	for i := 0; i < length; i++ {
		pair = append(pair, []int{start[i], end[i]})
	}
	sort.Slice(pair, func(i, j int) bool {
		return pair[i][1] < pair[j][1]
	})
	count, ansEnd := 1, pair[0][1]
	for i := 1; i < length; i++ {
		if pair[i][0] > ansEnd {
			count++
			ansEnd = pair[i][1]
		}
	}
	return count
}
