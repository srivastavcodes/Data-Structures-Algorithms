package main

import "sort"

func candyStore(candies []int, n int, k int) []int {
	sort.Ints(candies)
	mini, buyIdx, freeIdx := 0, 0, n-1

	for buyIdx <= freeIdx {
		mini = mini + candies[buyIdx]
		buyIdx++
		freeIdx = freeIdx - k
	}
	maxi, buyIdx, freeIdx := 0, n-1, 0

	for buyIdx >= freeIdx {
		maxi = maxi + candies[buyIdx]
		buyIdx--
		freeIdx = freeIdx + k
	}
	return []int{mini, maxi}
}
