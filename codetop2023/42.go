package codetop2023

func trap(height []int) int {
	sum := 0
	if len(height) <= 2 {
		return sum
	}

	for i := 1; i < len(height)-1; i++ {
		leftMax, rightMax := 0, 0
		// 寻找左边最高柱子
		for l := i - 1; l < i; l-- {
			leftMax = max(leftMax, height[l])
		}
		for r := i + 1; r < len(height); r++ {
			rightMax = max(rightMax, height[i+1])
		}

		minT := min(leftMax, rightMax)
		if minT > height[i] {
			sum += minT - height[i]
		}
	}
	return sum
}

func min(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
