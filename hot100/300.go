package hot100

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// 以nums[i]结尾的最长递增子序列的长度
	dp := make([]int, len(nums))
	// base case
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max300(dp[j]+1, dp[i])
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max300(res, dp[i])
	}
	return res
}

func max300(a, b int) int {
	if a > b {
		return a
	}
	return b
}
