package main

func main() {}

func trap(height []int) int {
	if len(height) < 1 {
		return 0
	}
	l, r := 0, len(height)-1
	result := 0
	lmax, rmax := height[l], height[r]

	for l < r {
		if lmax < rmax {
			l++
			lmax = max(lmax, height[l])
			result += lmax - height[l]
		} else {
			r--
			rmax = max(rmax, height[r])
			result += rmax - height[r]
		}
	}
	return result
}
