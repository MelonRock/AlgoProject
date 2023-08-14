package array

func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		dp[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j > i {
				dp[i] = min45(dp[j]+1, dp[i])
			}
		}
	}
	return dp[len(nums)-1]
}

func min45(a, b int) int {
	if a < b {
		return a
	}
	return b
}
