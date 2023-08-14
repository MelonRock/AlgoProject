package hot100

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if a > 0 {
			break
		}
		// 跳过相同数字
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			b, c := nums[l], nums[r]
			if a+b+c == 0 {
				ans := []int{a, b, c}
				res = append(res, ans)
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
