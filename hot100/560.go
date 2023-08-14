package hot100

// 边存边查看 map，如果 map 中存在 key 为「当前前缀和 - k」，
//说明这个之前出现的前缀和，满足「当前前缀和 - 该前缀和 == k」，
//它出现的次数，累加给 count。

func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if hash[preSum-k] > 0 {
			count += hash[preSum-k]
		}
		hash[preSum]++
	}
	return count
}
