package codetop2023

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// dp[i]是nums中以nums[i]结尾的最长递增子序列
	dp := make([]int, len(nums))
	// base case初始长度为1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := dp[0]
	for i := 0; i < len(nums); i++ {
		res = max(res, dp[i])
	}
	return res
}
