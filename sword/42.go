package sword

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max42(nums[i], dp[i-1]+nums[i])
	}
	res := nums[0]
	for i := 0; i < n; i++ {
		res = max42(res, dp[i])
	}
	return res
}

func max42(a, b int) int {
	if a > b {
		return a
	}
	return b
}
