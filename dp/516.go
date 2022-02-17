package dp

// 最长回文子序列 https://blog.csdn.net/jianghao1996/article/details/106258592
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
				//  s 的第 i 个字符和第 j 个字符不同的, i 从最后一个字符开始往前遍历，j 从 i + 1 开始往后遍历，这样可以保证每个子问题都已经算好了
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
