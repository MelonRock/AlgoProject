package sword

func cuttingRope(n int) int {
	// dp[n]表示长度为n的最大乘积
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			// dp[i]  维持不变
			// j * (i-j) 从j处剪下来，剪下来的部分是i-j, i-j不再剪
			// j*dp[i-j] 从j处剪一下，剪下来的部分是i-j, i-j继续剪
			dp[i] = max(max(j*(i-j), j*dp[i-j]), dp[i])
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
