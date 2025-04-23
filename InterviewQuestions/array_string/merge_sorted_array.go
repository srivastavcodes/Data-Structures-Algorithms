package main

import "fmt"

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	merge(nums1, 3, nums2, len(nums2))
	fmt.Printf("nums1: %v", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
