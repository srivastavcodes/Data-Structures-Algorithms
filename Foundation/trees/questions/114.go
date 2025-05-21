package main

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			l := curr.Left
			for l.Right != nil {
				l = l.Right
			}
			l.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}
