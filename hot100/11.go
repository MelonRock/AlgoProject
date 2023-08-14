package hot100

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		area := (right - left) * min11(height[left], height[right])
		max = max11(max, area)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}

func max11(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min11(a, b int) int {
	if a < b {
		return a
	}
	return b
}
