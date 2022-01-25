package dp

func longestPalindromeSubseq(s string) int {
	n := len(s)
	// f[i][j] 表示 s 的第 i 个字符到第 j 个字符组成的子串中，最长的回文序列长度是多少
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i] == nil {
				dp[i] = make([]int, n)
			}
			if i == j {
				dp[i][j] = 1
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
