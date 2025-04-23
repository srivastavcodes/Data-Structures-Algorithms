package main

import "fmt"

func main() {
	skip("", "baccdah")
	result := skipReturn("baaccdah")
	fmt.Println(result)
}

func skip(p, up string) {
	if len(up) == 0 {
		fmt.Println(p)
		return
	}
	ch := up[0]

	if ch == 'a' {
		skip(p, up[1:])
	} else {
		skip(p+string(ch), up[1:])
	}
}

func skipReturn(up string) string {
	if len(up) == 0 {
		return ""
	}
	ch := up[0]

	if ch == 'a' {
		return skipReturn(up[1:])
	} else {
		return string(ch) + skipReturn(up[1:])
	}
}
