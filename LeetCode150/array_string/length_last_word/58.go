package main

func lengthOfLastWord(str string) int {
	i, length := len(str)-1, 0

	for str[i] == ' ' {
		i--
	}
	for i > 0 && str[i] != ' ' {
		length++
	}
	return length
}
