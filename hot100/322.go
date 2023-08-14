package hot100

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; i < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min322(dp[i-1], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return 0
	}
	return dp[amount]
}

func min322(a, b int) int {
	if a < b {
		return a
	}
	return b
}
