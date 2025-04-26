package main

func trap(height []int) int {
	lmax, rmax, result, l, r := 0, 0, 0, 0, len(height)-1
	for l < r {
		if height[l] <= height[r] {
			if lmax > height[l] {
				result += lmax - height[l]
			} else {
				lmax = height[l]
			}
			l++
		} else {
			if rmax > height[r] {
				result += rmax - height[r]
			} else {
				rmax = height[r]
			}
			r--
		}
	}
	return result
}
