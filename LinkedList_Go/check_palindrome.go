package main

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var secondHalf *ListNode
	if fast == nil {
		secondHalf = slow
	} else {
		secondHalf = slow.Next
	}
	secondHalf = reverseList(secondHalf)

	result := true
	first := head
	second := secondHalf

	for second != nil {
		if first.Val != second.Val {
			result = false
			break
		}
		first = first.Next
		second = second.Next
	}
	reverseList(secondHalf)
	return result
}
