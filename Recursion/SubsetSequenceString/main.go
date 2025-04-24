package main

import "fmt"

func main() {
	perm := permutations("", "abc")
	fmt.Println(perm)

	digits := []int{9, 0, 3, 5, 7}
	subs := subsets(digits)
	fmt.Println(subs)
}
