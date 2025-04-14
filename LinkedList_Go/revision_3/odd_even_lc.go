package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var odd = head
	var even, evenHead = head.Next, head.Next

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next

		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
