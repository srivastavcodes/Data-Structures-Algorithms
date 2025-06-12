package main

import "fmt"

func main() {
	strs := "AAAAAAAAAAA"

	repeat := findRepeatedDnaSequences(strs)
	fmt.Println(repeat)
}
