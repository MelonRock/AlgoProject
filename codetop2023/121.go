package codetop2023

func maxProfit(price []int) int {
	n := len(price)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	// base case
	dp[0][0] = 0
	dp[0][1] = -price[0]

	for i := 1; i < n; i++ {
		// 没有持有，前一天没有持有或者当天卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i])
		// 持有，前一天持有或者当天买入
		dp[i][1] = max(dp[i-1][1], -price[i])
	}
	return dp[n-1][0]
}
