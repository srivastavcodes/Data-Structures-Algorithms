package main

type Node struct {
	Left  *Node
	Val   int
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	leftMost := root

	for leftMost.Left != nil {
		current := leftMost

		for current != nil {
			current.Left.Next = current.Right

			if current.Next != nil {
				current.Right.Next = current.Next.Left
			}
			current = current.Next
		}
		leftMost = leftMost.Left
	}
	return root
}
