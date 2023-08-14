package array

// nums[i] = i+1
func firstMissingPositive(nums []int) int {
	n := len(nums)
	temp := make([]int, n+2)

	for i := 0; i < n; i++ {
		if nums[i] >= 0 && nums[i] <= n+1 {
			temp[nums[i]] = nums[i]
		}
	}
	for i := 1; i <= n+1; i++ {
		if temp[i] == 0 {
			return i
		}
	}
	return 0
}
