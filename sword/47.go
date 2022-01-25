package sword

func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid)+1)
	for i := 0; i <= len(grid); i++ {
		dp[i] = make([]int, len(grid[0])+1)
	}

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = grid[i-1][j-1] + max(dp[i][j-1], dp[i-1][j])
		}
	}
	return dp[len(grid)][len(grid[0])]
}
