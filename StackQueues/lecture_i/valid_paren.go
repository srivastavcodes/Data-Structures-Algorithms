package main

var st = "(){[]}"

func isValid(str string) bool {
	if len(str) < 2 {
		return false
	}
	record := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune

	for _, s := range str {
		if _, ok := record[s]; ok {
			if len(stack) != 0 && stack[len(stack)-1] == record[s] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s)
		}
	}
	return len(stack) == 0
}
