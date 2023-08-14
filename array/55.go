package array

func canJump(nums []int) bool {
	cover := 0
	n := len(nums) - 1
	// 每次与覆盖值比较
	for i := 0; i <= cover; i++ {
		// 每走一步都将cover更新为最大值
		cover = max55(i+nums[i], cover)
		if cover >= n {
			return true
		}
	}
	return false
}

func max55(a, b int) int {
	if a > b {
		return a
	}
	return b
}
