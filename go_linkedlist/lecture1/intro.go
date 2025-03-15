package lecture1

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConstructLL() *ListNode {
	vals := []int{2, 3, 4, 5, 6, 7}
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
