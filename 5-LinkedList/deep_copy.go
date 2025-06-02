package main

func copyRandomList(head *Node) *Node {
	insertCopyInBetween(head)
	connectRandomPointers(head)
	return getDeepCopyList(head)
}

func getDeepCopyList(head *Node) *Node {
	var temp = head
	dummyNode := &Node{Val: -1, Next: nil}
	res := dummyNode

	for temp != nil {
		res.Next = temp.Next
		res = res.Next

		temp.Next = temp.Next.Next
		temp = temp.Next
	}
	return dummyNode.Next
}

func insertCopyInBetween(head *Node) {
	var temp = head
	for temp != nil {
		var nextElement = temp.Next
		copied := &Node{Val: temp.Val}

		copied.Next = nextElement
		temp.Next = copied
		temp = nextElement
	}
}

func connectRandomPointers(head *Node) {
	var temp = head
	for temp != nil {
		var copyNode = temp.Next
		if temp.Random != nil {
			copyNode.Random = temp.Random.Next
		} else {
			copyNode.Random = nil
		}
		temp = temp.Next.Next
	}
}
