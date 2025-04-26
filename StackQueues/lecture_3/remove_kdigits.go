package main

func removeKdigits(num string, k int) string {
	if k > len(num) {
		return "0"
	}
	stack := make([]byte, 0, len(num))

	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && k > 0 && (stack[len(stack)-1]) > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	index := 0
	for index < len(stack) && stack[index] == '0' {
		index++
	}
	if len(stack) == index {
		return "0"
	}
	return string(stack[index:])
}
