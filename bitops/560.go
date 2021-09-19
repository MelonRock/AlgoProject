package bitops

func subarraySum(nums []int, k int) int {
	n := len(nums)
	sums := make([]int, n+1)
	ans := 0
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if sums[i]-sums[j] == k {
				ans++
			}
		}
	}
	return ans
}
