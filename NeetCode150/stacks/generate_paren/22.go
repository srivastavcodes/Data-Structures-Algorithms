package main

import "fmt"

func main() {
	n := 3
	gen := generateParenthesis(n)
	fmt.Println(gen)
}

func generateParenthesis(count int) []string {
	return getParen("", count, 0, 0)
}

func getParen(str string, count int, open, closed int) []string {
	result := make([]string, 0)

	if open == count && closed == count {
		return []string{str}
	}
	if open < count {
		result = append(result, getParen(str+"(", count, open+1, closed)...)
	}
	if closed < open {
		result = append(result, getParen(str+")", count, open, closed+1)...)
	}
	return result
}
