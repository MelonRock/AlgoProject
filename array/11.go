package array

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		// 求面积
		area := (right - left) * min11(height[left], height[right])
		// 更新最大值
		ans = max11(ans, area)
		// 收缩移动短的那一根
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

func max11(a, b int) int {
	if a > b {
		return a
	}
	return b
}
