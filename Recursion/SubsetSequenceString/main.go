package main

import "fmt"

func main() {
	skip('c', "", "baccdah")
	result := skipReturn('a', "baaccdah")
	skippedString := skipString("apple", "badappleiate")
	fmt.Println(result)
	fmt.Println(skippedString)

	var resList []string
	subseq("", "123", &resList)
	fmt.Println(resList)

	var resList2 []string
	resList2 = subseq2("", "123")
	fmt.Println(resList2)
}
