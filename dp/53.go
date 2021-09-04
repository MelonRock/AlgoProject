package dp

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	// base cese
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		// 要么自成一派，要么和前面的子数组合并
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}
	res := dp[0]
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
