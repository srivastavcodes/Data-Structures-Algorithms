package main

import (
	"fmt"

	"go_linkedlist/lecture1"
)

func main() {
	ll := lecture1.ConstructLL()
	fmt.Printf("%v", lecture1.Length(ll))
}
