package codetop2023

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	return nums1
}
