package main

import "fmt"

type Heap struct {
	vals []int
	size int
}

func (h *Heap) insert(val int) {
	h.size++
	if len(h.vals) < h.size {
		h.vals = append(h.vals, val)
	} else {
		h.vals[h.size-1] = val
	}

	idx := h.size - 1
	for idx > 0 {
		parent := (idx - 1) / 2
		if h.vals[parent] < h.vals[idx] {
			h.vals[parent], h.vals[idx] = h.vals[idx], h.vals[parent]
			idx = parent
		} else {
			break
		}
	}
}

func (h *Heap) deleteNode() {
	if h.size == 0 {
		fmt.Printf("Nothing to delete")
		return
	}
	h.vals[0] = h.vals[h.size-1]
	h.size--
	h.vals = h.vals[:h.size]

	idx := 0
	for {
		leftIdx := 2*idx + 1
		currIdx := idx
		rightIdx := 2*idx + 2

		if leftIdx < h.size && h.vals[currIdx] < h.vals[leftIdx] {
			currIdx = leftIdx
		}
		if rightIdx < h.size && h.vals[currIdx] < h.vals[rightIdx] {
			currIdx = rightIdx
		}
		if currIdx == idx {
			break
		}
		h.vals[idx], h.vals[currIdx] = h.vals[currIdx], h.vals[idx]
		idx = currIdx
	}
}

func (h *Heap) printHeap() {
	fmt.Printf("heap :: \t%d\n", h.vals)
}
