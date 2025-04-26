package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var nge []int
	for _, num := range nums1 {
		var stack []int
		var result bool
		for j := len(nums2) - 1; num != nums2[j]; j-- {
			if num < nums2[j] {
				result = true
				stack = append(stack, nums2[j])
			}
		}
		if !result {
			nge = append(nge, -1)
		} else {
			nge = append(nge, stack[len(stack)-1])
		}
	}
	return nge
}

func nextGreaterElementOptimal(nums1 []int, nums2 []int) []int {
	record := make(map[int]int)
	var stack []int

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			record[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i, num := range nums1 {
		if val, ok := record[num]; ok {
			nums1[i] = val
		} else {
			nums1[i] = -1
		}
	}
	return nums1
}
