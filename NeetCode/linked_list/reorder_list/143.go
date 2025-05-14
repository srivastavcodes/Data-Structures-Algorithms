package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := slow.Next
	slow.Next = nil
	var prev *ListNode

	for second != nil {
		second.Next, prev, second = prev, second, second.Next
	}
	first, second := head, prev
	for second != nil {
		first.Next, second.Next, first, second = second, first.Next, first.Next, second.Next
	}
}
