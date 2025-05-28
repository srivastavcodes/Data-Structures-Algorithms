package main

import "fmt"

func main() {
	strs := "forxxorfxdofr"
	pattern := "for"

	count := countOccurrence(pattern, strs)
	fmt.Println(count)
}
