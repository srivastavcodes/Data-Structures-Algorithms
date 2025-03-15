package main

import (
	"fmt"

	"go_linkedlist/lecture1"
)

func main() {
	vals := []int{1, 2, 3, 4, 5}
	ll := lecture1.ConstructLL(vals)
	fmt.Printf("%v", lecture1.Length(ll))
	fmt.Printf("\n%v", lecture1.SearchInLinkedList(ll, 11))
}
