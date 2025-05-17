package main

import "fmt"

func main() {
	nums := 3245
	roman := intToRoman(nums)
	fmt.Println(roman)
}

func intToRoman(num int) string {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := ""

	for i, val := range values {
		for num >= val {
			roman += symbols[i]
			num -= val
		}
	}
	return roman
}
