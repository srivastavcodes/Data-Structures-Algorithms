package main

import "fmt"

func main() {
	fmt.Println("Greedy Problem Begins--------------------")

	input := []int{1, 2, 3}
	count := jump(input)
	fmt.Println(count)
}
