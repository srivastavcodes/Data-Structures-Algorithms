package main

import (
	"fmt"
	"strings"
)

func main() {
	skip('c', "", "baccdah")
	result := skipReturn('a', "baaccdah")
	skippedString := skipString("apple", "badappleiate")
	fmt.Println(result)
	fmt.Println(skippedString)
}

func skip(char uint8, p, up string) {
	if len(up) == 0 {
		fmt.Println(p)
		return
	}
	ch := up[0]

	if ch == char {
		skip(char, p, up[1:])
	} else {
		skip(char, p+string(ch), up[1:])
	}
}

func skipReturn(char uint8, up string) string {
	if len(up) == 0 {
		return ""
	}
	ch := up[0]

	if ch == char {
		return skipReturn(char, up[1:])
	} else {
		return string(ch) + skipReturn(char, up[1:])
	}
}

func skipString(value, str string) string {
	if len(str) == 0 {
		return ""
	}
	if strings.HasPrefix(str, value) {
		return skipString(value, str[len(value):])
	} else {
		return string(str[0]) + skipString(value, str[1:])
	}
}
