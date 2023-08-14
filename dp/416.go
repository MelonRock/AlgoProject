package dp

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	// dp[i][j] = x 表示，对于前 i 个物品，
	// 当前背包的容量为 j 时，若 x 为 true，则说明可以恰好将背包装满，
	// 若 x 为 false，则说明不能恰好将背包装满。
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ {
			//
			if j-nums[i-1] >= 0 {
				// 装入或不装入背包
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			} else {
				// 背包容量不足，不能装入第 i 个物品
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum]
}
