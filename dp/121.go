package dp

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		// 没有持有
		dp[i][0] = max5(dp[i-1][0], dp[i-1][1]+prices[i])
		// 持有
		dp[i][1] = max5(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

func max5(a, b int) int {
	if a > b {
		return a
	}
	return b
}
