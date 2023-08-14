package hot100

func longestConsecutive(nums []int) int {
	// 创建一个哈希集合，用于存储数组中的数字
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	// 记录最长连续序列的长度
	longestStreak := 0
	// 遍历哈希集合中的每个数字
	for num := range numSet {
		// 如果当前数字的前一个数字不在哈希集合中，则当前数字是连续序列的开头
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			// 继续递增当前数字，直到连续序列结束
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			// 更新最长连续序列的长度
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	// 返回最长连续序列的长度
	return longestStreak
}
