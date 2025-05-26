package main

func jump(nums []int) int {
	count, l, r := 0, 0, 0

	for r < len(nums)-1 {
		leap := 0
		for i := l; i <= r; i++ {
			leap = max(leap, i+nums[i])
		}
		l++
		count++
		r = leap
	}
	return count
}
