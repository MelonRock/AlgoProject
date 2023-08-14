package hot100

func search(nums []int, target int) int {
	// 先根据 nums[mid] 与 nums[lo] 的关系判断 mid 是在左段还是右段，
	//接下来再判断 target 是在 mid 的左边还是右边，从而来调整左右边界 lo 和 hi。
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		// mid在左段
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[lo] {
			// 判断target位置
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			// mid在右段
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
