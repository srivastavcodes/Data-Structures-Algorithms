package main

import (
	"fmt"
	"math"
)

func main() {
	bp := []int{7, 2, 5, 10, 8}
	stdnts := 2
	result := splitArray(bp, stdnts)
	fmt.Print(result)
}

func splitArray(books []int, stdnts int) int {
	if stdnts > len(books) {
		return -1
	}
	low, high := math.MinInt, 0
	for _, pgs := range books {
		low = max(low, pgs)
		high += pgs
	}
	for low <= high {
		cntr := (low + high) / 2
		if valid(books, cntr, stdnts) {
			high = cntr - 1
		} else {
			low = cntr + 1
		}
	}
	return low
}

func valid(books []int, cntr int, stdnt int) bool {
	ttlPgs, count := 0, 1
	for _, pgs := range books {
		if ttlPgs+pgs <= cntr {
			ttlPgs += pgs
		} else {
			count++
			ttlPgs = pgs
		}
	}
	return count <= stdnt
}
