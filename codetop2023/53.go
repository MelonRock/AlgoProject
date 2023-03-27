package codetop2023

func maxSubArray(nums []int) int {
	n := len(nums)
	// dp[i]表示以nums[i]结尾的 连续 子数组的最大和
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		// 要么自成一派，要么与前面的子数组相连
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	res := nums[0]
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}
