package array

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	sortedNum2 := [][2]int{}
	for index, value := range nums2 {
		sortedNum2 = append(sortedNum2, [2]int{value, index})
	}
	// 升序排列
	sort.Slice(sortedNum2, func(i, j int) bool {
		return sortedNum2[i][0] < sortedNum2[j][0]
	})
	// 升序排列
	sort.Ints(nums1)
	res := make([]int, len(nums1))
	left := 0 // nums1 最小值，记录用了多少个左边的值
	for i := len(sortedNum2) - 1; i >= 0; i-- {
		//
		if nums1[i+left] > sortedNum2[i][0] {
			res[sortedNum2[i][1]] = nums1[i+left]
		} else {
			res[sortedNum2[i][1]] = nums1[left]
			left++
		}
	}
	return res
}
