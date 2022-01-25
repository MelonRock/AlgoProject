package codetop

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// dp[i]：表示以 nums[i] 结尾 的 连续 子数组的最大和。
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max1(dp[i-1]+nums[i], nums[i])
	}
	res := nums[0]
	for i := 0; i < n; i++ {
		res = max1(res, dp[i])
	}
	return res
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
