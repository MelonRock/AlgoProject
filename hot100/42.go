package hot100

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	size := len(height)
	maxLeft := make([]int, len(height))
	maxLeft[0] = height[0]
	for i := 1; i < size; i++ {
		maxLeft[i] = max42(height[i], maxLeft[i-1])
	}
	maxRight := make([]int, len(height))
	maxRight[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		maxRight[i] = max42(height[i], maxRight[i+1])
	}
	sum := 0
	for i := 0; i < size; i++ {
		count := min42(maxLeft[i], maxRight[i]) - height[i]
		if count > 0 {
			sum += count
		}
	}
	return sum
}

func max42(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min42(a, b int) int {
	if a < b {
		return a
	}
	return b
}
