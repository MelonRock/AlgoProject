package hot100

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			// base case
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		// 没有持有
		dp[i][0] = max121(dp[i-1][0], prices[i]+dp[i-1][1])
		// 持有
		dp[i][1] = max121(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

func max121(a, b int) int {
	if a > b {
		return a
	}
	return b
}
