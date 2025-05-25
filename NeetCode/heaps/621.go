package main

import "github.com/emirpasic/gods/queues/priorityqueue"

func leastInterval(tasks []byte, n int) int {
	record := make([]int, 26)

	for _, task := range tasks {
		record[task-'A']++
	}
	var time int
	maxHeap := priorityqueue.NewWith(func(a, b interface{}) int {
		return b.(int) - a.(int)
	})
	for i := 0; i < 26; i++ {
		if record[i] > 0 {
			maxHeap.Enqueue(record[i])
		}
	}
	for !maxHeap.Empty() {
		var temp []int
		for i := 1; i <= n+1; i++ {
			if !maxHeap.Empty() {
				freqIF, _ := maxHeap.Dequeue()
				freq := freqIF.(int)
				freq--
				temp = append(temp, freq)
			}
		}
		for _, freq := range temp {
			if freq > 0 {
				maxHeap.Enqueue(freq)
			}
		}
		if maxHeap.Empty() {
			time += len(temp)
		} else {
			time += n + 1
		}
	}
	return time
}
