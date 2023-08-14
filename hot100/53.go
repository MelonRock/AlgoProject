package hot100

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	// dp[i] 以第i个数结尾的最大连续子数组和的最大值
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 当前数字与前面的合成一组或者自己自成一组，取较大值
		dp[i] = max53(dp[i-1]+nums[i], nums[i])
	}
	res := dp[0]
	for i := 0; i < len(dp); i++ {
		res = max53(res, dp[i])
	}
	return res
}

func max53(a, b int) int {
	if a > b {
		return a
	}
	return b
}
