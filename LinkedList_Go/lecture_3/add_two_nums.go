package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: nil}
	curr := dummyHead
	temp1, temp2 := l1, l2
	var carry = 0
	for temp1 != nil || temp2 != nil {
		sum := carry
		if temp1 != nil {
			sum += temp1.Val
		}
		if temp2 != nil {
			sum += temp2.Val
		}
		carry = sum / 10

		curr.Next = &ListNode{Val: sum % 10, Next: nil}
		curr = curr.Next
		if temp1 != nil {
			temp1 = temp1.Next
		}
		if temp2 != nil {
			temp2 = temp2.Next
		}
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry, Next: nil}
	}
	return dummyHead.Next
}
