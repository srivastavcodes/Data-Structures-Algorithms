package main

func removeDuplicates(strs string) string {
	var stack []byte

	for _, str := range strs {
		if len(stack) > 0 && stack[len(stack)-1] == byte(str) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(str))
		}
	}
	return string(stack)
}
