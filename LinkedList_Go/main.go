package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

// PrintLinkedList prints all values in the linked list
func PrintLinkedList(head *ListNode) {
	if head == nil {
		fmt.Println("Empty list")
		return
	}

	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	list := CreateLinkedList([]int{1, 2, 3, 4, 5})

	fmt.Println("Original linked list:")
	PrintLinkedList(list)

	deleteNode(list.Next.Next)
	fmt.Printf("List after deletion:\n")
	PrintLinkedList(list)
}
