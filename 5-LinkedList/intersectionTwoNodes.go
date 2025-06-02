package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var temp1 = headA
	var temp2 = headB
	for temp1 != temp2 {
		temp1 = temp1.Next
		temp2 = temp2.Next
		if temp1 == temp2 {
			return temp1
		}
		if temp1 == nil {
			temp1 = headB
		}
		if temp2 == nil {
			temp2 = headA
		}
	}
	return temp1
}
