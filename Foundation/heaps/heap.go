package main

import "fmt"

type Heap struct {
	vals []int
	size int
}

func (h *Heap) insert(val int) {
	h.size++
	if h.size >= len(h.vals) {
		h.vals = append(h.vals, val)
	} else {
		h.vals[h.size] = val
	}

	idx := h.size
	for idx > 1 {
		parent := idx / 2
		if h.vals[parent] < h.vals[idx] {
			h.vals[parent], h.vals[idx] = h.vals[idx], h.vals[parent]
			idx = parent
		} else {
			return
		}
	}
}

func (h *Heap) deleteNode() {
	if h.size == 0 {
		fmt.Printf("Nothing to delete")
		return
	}
	h.vals[1] = h.vals[h.size]
	h.size--

	// take root node to it correct position
	idx := 1
	for idx < h.size {
		leftIdx := 2 * idx
		rightIdx := 2*idx + 1

		if leftIdx < h.size && h.vals[idx] < h.vals[leftIdx] {
			h.vals[leftIdx], h.vals[idx] = h.vals[idx], h.vals[leftIdx]
			idx = leftIdx
		} else if rightIdx < h.size && h.vals[idx] < h.vals[rightIdx] {
			h.vals[rightIdx], h.vals[idx] = h.vals[idx], h.vals[rightIdx]
			idx = rightIdx
		} else {
			return
		}
	}
}

func (h *Heap) printHeap() {
	fmt.Printf("heap :: \t%d\n", h.vals)
}

func main() {
	heap := Heap{vals: make([]int, 1, 10)}
	heap.insert(54)
	heap.insert(53)
	heap.insert(55)
	heap.insert(52)
	heap.insert(50)

	heap.printHeap()

	vals := []int{-1, 54, 53, 55, 52, 50}
	fmt.Printf("vals :: \t%d", vals)
}
