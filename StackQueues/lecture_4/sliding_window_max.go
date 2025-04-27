package lecture_4

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums))
	slide := make([]int, 0, k)

	for i, left := range nums {
		for len(slide) > 0 && nums[slide[len(slide)-1]] < left {
			slide = slide[:len(slide)-1]
		}
		if len(slide) > 0 && slide[0] == i-k {
			slide = slide[1:]
		}
		slide = append(slide, i)
		if i >= k-1 {
			res = append(res, nums[slide[0]])
		}
	}
	return res
}
