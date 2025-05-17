package roman_to_integer

func romanToInt(str string) int {
	roman := map[string]int{
		"I": 1, "V": 5, "X": 10,
		"L": 50, "C": 100, "D": 500, "M": 1000,
	}
	result := 0
	for i := range str {
		if i+1 < len(str) && roman[string(str[i])] < roman[string(str[i+1])] {
			result -= roman[string(str[i])]
		} else {
			result += roman[string(str[i])]
		}
	}
	return result
}
