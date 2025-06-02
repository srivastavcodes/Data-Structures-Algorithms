package main

func middleNode(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var back *ListNode
	var curr = head
	for curr != nil {
		next := curr.Next
		curr.Next = back
		back = curr
		curr = next
	}
	return back
}
