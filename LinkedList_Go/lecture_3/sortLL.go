package main

func findMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var slow, fast = head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	var dummyNode = &ListNode{}
	var temp = dummyNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1
			temp = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			temp = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		temp.Next = list1
	} else {
		temp.Next = list2
	}
	return dummyNode.Next
}

func sortLL(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	middle := findMiddle(head)

	right := middle.Next
	middle.Next = nil
	left := head

	left = sortLL(left)
	right = sortLL(right)
	return mergeTwoLists(left, right)
}
