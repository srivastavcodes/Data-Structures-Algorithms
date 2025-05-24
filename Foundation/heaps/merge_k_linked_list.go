package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
)

type Node struct {
	data int
	next *Node
}

func mergeKLists(lists []*Node) *Node {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		val1, val2 := a.(*Node), b.(*Node)
		switch {
		case val1.data < val2.data:
			return -1
		case val1.data > val2.data:
			return 1
		default:
			return 0
		}
	})
	if len(lists) == 0 {
		return nil
	}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			pq.Enqueue(lists[i])
		}
	}
	node := &Node{}
	tail := node

	for !pq.Empty() {
		currIF, _ := pq.Dequeue()
		curr := currIF.(*Node)

		tail.next = curr
		tail = tail.next

		if curr.next != nil {
			pq.Enqueue(curr.next)
		}
	}
	return node.next
}
