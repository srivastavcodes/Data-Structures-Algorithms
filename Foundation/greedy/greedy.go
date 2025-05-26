package main

import "fmt"

func main() {
	input := []int{3, 2, 1, 4}

	minmax := candyStore(input, len(input), 2)
	fmt.Println(minmax)
}
