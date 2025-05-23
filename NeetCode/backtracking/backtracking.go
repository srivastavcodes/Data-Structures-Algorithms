package main

import "fmt"

func main() {
	digits := []int{4, 4, 4, 1, 4}

	// output := subsets(digits)
	// fmt.Println(output)

	// comb := combinationSum2(digits, 8)
	// fmt.Println(comb)

	// perm := permute(digits)
	// fmt.Println(perm)

	subsetII := subsetsWithDup(digits)
	fmt.Println(subsetII)
}
