package codetop2023

import "sort"

func threeSum(nums []int) [][]int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left int, right int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	if left >= right {
		return res
	}
	for i := 0; i < len(nums)-2; i++ {
		// 取第一个数
		a := nums[i]
		if a > 0 {
			return res
		}
		// 跳过相同的数字
		if i > 0 && a == nums[i-1] {
			continue
		}
		// 其余两个数的下标
		l, r := i+1, len(nums)-1
		for l < r {
			b, c := nums[l], nums[r]
			if a+b+c == 0 {
				res = append(res, []int{a, b, c})
				// 跳过相同的数字
				for l < r && b == nums[l] {
					l++
				}
				for l < r && c == nums[r] {
					r--
				}
			} else if a+b+c > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
