package dp

// dp[i] 表示以 nums[i] 这个数结尾的最长递增子序列的长度。
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1) // nums[i]可以接在 nums[j] 之后
			}
		}
	}
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
