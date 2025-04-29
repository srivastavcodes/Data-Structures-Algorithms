package main

import "sort"

func main() {}

func threeSum(digits []int) [][]int {
	var result [][]int
	sort.Ints(digits)

	for i := 0; i < len(digits); i++ {
		a := digits[i]
		if a > 0 {
			break
		}
		if i > 0 && a == digits[i-1] {
			continue
		}
		left, right := i+1, len(digits)-1
		for left < right {
			sum := a + digits[left] + digits[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{a, digits[left], digits[right]})
				left++
				right--
				for left < right && digits[left] == digits[left-1] {
					left++
				}
			}
		}
	}
	return result
}
