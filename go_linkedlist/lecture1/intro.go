package lecture1

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConstructLL(vals []int) *ListNode {
	head := &ListNode{Val: vals[0], Next: nil}
	tail := head
	for i := 1; i < len(vals); i++ {
		temp := &ListNode{Val: vals[i], Next: nil}
		tail.Next = temp
		tail = temp
	}
	return head
}

func Length(node *ListNode) int {
	lstNode, count := node, 0
	for lstNode != nil {
		count++
		lstNode = lstNode.Next
	}
	return count
}

func SearchInLinkedList(node *ListNode, val int) bool {
	lstNode := node
	for lstNode != nil {
		if lstNode.Val == val {
			return true
		}
		lstNode = lstNode.Next
	}
	return false
}
