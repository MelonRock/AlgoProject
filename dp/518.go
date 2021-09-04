package dp

func change(amount int, coins []int) int {
	n := len(coins)
	// 若只使用 coins 中的前 i 个硬币的面值，若想凑出金额 j，有 dp[i][j] 种凑法。
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				// 不使用和使用第i枚硬币
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][amount]
}
