package codetop

func lengthOfLIS(nums []int) int {
	// dp[i]表示以nums[i]结尾的最长递增子序列
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max2(dp[i], dp[j]+1)
			}
		}
	}
	res := dp[0]
	for i := 1; i < len(dp); i++ {
		res = max2(res, dp[i])
	}
	return res
}

func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
