package hot100

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]

		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				if n1+n2+n3+n4 == target {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && n3 == nums[l] {
						l++
					}
					for l < r && n4 == nums[r] {
						r--
					}
				} else if n1+n2+n3+n4 < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
