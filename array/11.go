package array

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min11(height[left], height[right])
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}

func min11(a, b int) int {
	if a > b {
		return b
	}
	return a
}
