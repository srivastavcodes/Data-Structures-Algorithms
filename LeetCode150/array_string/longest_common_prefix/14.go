package main

import "fmt"

func main() {
	str := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(str)
	fmt.Println(prefix)
}

func longestCommonPrefix(str []string) string {
	result := ""
	for i := range str[0] {
		for _, s := range str {
			if i == len(s) || s[i] != str[0][i] {
				return result
			}
		}
		result += string(str[0][i])
	}
	return result
}
