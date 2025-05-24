package main

import "fmt"

func main() {
	dict := Constructor()
	input := []string{"bad", "dad", "mad", "badminton"}
	for _, str := range input {
		dict.AddWord(str)
	}
	fmt.Println(dict)

	search := []string{".ad", "b.."}
	for _, s := range search {
		res := dict.Search(s)
		fmt.Println(res)
	}
}
