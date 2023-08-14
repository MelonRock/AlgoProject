package codetop2023

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] >= 0 {
				// 装入或者不装
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			} else {
				// 背包容量不够
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum]
}
