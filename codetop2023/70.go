package codetop2023

func climbStairs(n int) int {
	// 爬第n阶楼梯的方法数量等于2部分之和
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
