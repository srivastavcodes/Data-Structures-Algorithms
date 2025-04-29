package main

func main() {}

func twoSum(digits []int, k int) []int {
	left, right := 0, len(digits)-1

	for left < right {
		sum := digits[left] + digits[right]
		if sum > k {
			right--
		} else if sum < k {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}
