package dp

func minDistance1(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max3(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	lcs := dp[len(text1)][len(text2)]
	return len(text1) - lcs + len(text2) - lcs
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
