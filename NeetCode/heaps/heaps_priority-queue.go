package main

import "fmt"

func main() {
	input := []int{3, 2, 1, 5, 6, 4}

	largest := findKthLargest(input, 2)
	fmt.Println(largest)
}
