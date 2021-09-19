package highFreq

func trap(height []int) int {
	n := len(height)
	res := 0
	for i := 0; i < n-1; i++ {
		l_max, r_max := 0, 0
		// 找右边最高
		for j := i; j < n; j++ {
			r_max = max(r_max, height[j])
		}
		// 找左边最高
		for j := i; j >= 0; j-- {
			l_max = max(l_max, height[j])
		}
		res += min(r_max, l_max) - height[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
