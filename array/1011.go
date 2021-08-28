package array

// 二分搜索，类似吃香蕉
func shipWithinDays(weights []int, D int) int {
	maxNum, sum := 0, 0
	for _, num := range weights {
		if num > maxNum {
			maxNum = num
		}
		sum += num
	}

	if D == 1 {
		return sum
	}
	left, right := maxNum, sum
	for left <= right {
		mid := left + (right-left)>>1
		if canShip(mid, D, weights) {
			// 需要让 f(x) 的返回值大一些, 搜索左侧边界
			right = mid - 1
		} else {
			// 需要让 f(x) 的返回值小一些
			left = mid + 1
		}
	}
	return left
}

func canShip(mid, D int, nums []int) bool {
	count := 1
	has := 0
	for i := range nums {
		if has+nums[i] <= mid {
			has += nums[i]
		} else {
			has = nums[i]
			count++
		}
	}
	return count <= D
}
