package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type MedianFinder struct {
	leftMaxHeap  *priorityqueue.Queue
	rightMinHeap *priorityqueue.Queue
}

func NewMedianFinder() MedianFinder {
	left := priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(int) - a.(int)
	})
	right := priorityqueue.NewWith(utils.IntComparator)
	return MedianFinder{
		leftMaxHeap:  left,
		rightMinHeap: right,
	}
}

func (mf *MedianFinder) AddNum(num int) {
	top, _ := mf.leftMaxHeap.Peek()
	if mf.leftMaxHeap.Empty() || num < top.(int) {
		mf.leftMaxHeap.Enqueue(num)
	} else {
		mf.rightMinHeap.Enqueue(num)
	}
	// maintain leftMaxHeap's size() > rightMinHeap's size() or equal()
	if abs(mf.leftMaxHeap.Size()-mf.rightMinHeap.Size()) > 1 {
		left, _ := mf.leftMaxHeap.Dequeue()
		mf.rightMinHeap.Enqueue(left)
	}
	if mf.leftMaxHeap.Size() < mf.rightMinHeap.Size() {
		right, _ := mf.rightMinHeap.Dequeue()
		mf.leftMaxHeap.Enqueue(right)
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.leftMaxHeap.Size() == mf.rightMinHeap.Size() {
		leftIF, _ := mf.leftMaxHeap.Peek()
		left := leftIF.(int)

		rightIF, _ := mf.rightMinHeap.Peek()
		right := rightIF.(int)

		return float64(left+right) / 2
	}
	leftIF, _ := mf.leftMaxHeap.Peek()
	return float64(leftIF.(int))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
