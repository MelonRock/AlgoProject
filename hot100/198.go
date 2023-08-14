package hot100

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1) // dp[i]表示打劫到第i个能偷到的最大金额
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max198(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}

func max198(a, b int) int {
	if a > b {
		return a
	}
	return b
}
