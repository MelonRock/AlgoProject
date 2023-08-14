package codetop2023

func searchRange(nums []int, target int) []int {
	// 两次二分查找，查找两次边界
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l, r := 0, len(nums)-1
	// 找左边界
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target { // 可以把右边裁剪掉
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[r] != target {
		return []int{-1, -1}
	}
	L := r
	// 找右边界
	l, r = 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] <= target { // 可以把左边裁剪掉
			l = mid
		} else {
			r = mid - 1
		}
	}
	return []int{L, r}
}
