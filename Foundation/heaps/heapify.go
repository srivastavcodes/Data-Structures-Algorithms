package main

import "fmt"

func heapify(vals []int, size, idx int, str string) {
	if str == "" {
		fmt.Printf("cannot be empty")
		return
	}
	if str == "max" {
		for i := len(vals) / 2; i > 0; i-- {
			heapifyMax(vals, len(vals), i)
		}
	} else if str == "min" {
		for i := len(vals)/2 - 1; i >= 0; i-- {
			heapifyMax(vals, len(vals), i)
		}
	}
}

// zero-based indexing
func heapifyMax(vals []int, size int, idx int) {
	largest := idx
	left, right := 2*idx, 2*idx+1

	if left < size && vals[largest] < vals[left] {
		largest = left
	} else if right < size && vals[largest] < vals[right] {
		largest = right
	}
	if largest != idx {
		vals[largest], vals[idx] = vals[idx], vals[largest]
		heapifyMax(vals, size, largest)
	}
}

// one-based indexing
func heapifyMin(vals []int, size int, idx int) {
	smallest := idx
	left, right := 2*idx+1, 2*idx+2

	if left < size && vals[smallest] > vals[left] {
		smallest = left
	} else if right < size && vals[smallest] > vals[right] {
		smallest = right
	}
	if smallest != idx {
		vals[smallest], vals[idx] = vals[idx], vals[smallest]
		heapifyMin(vals, size, smallest)
	}
}
