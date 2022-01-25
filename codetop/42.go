package codetop

func trap(height []int) int {
	sum := 0
	if len(height) <= 2 {
		return sum
	}

	for i := 1; i < len(height)-1; i++ {
		leftMax, rightMax := 0, 0
		// 寻找左边最高柱子
		for l := i - 1; l >= 0; l-- {
			leftMax = maxT(leftMax, height[l])
		}
		// 寻找右边最高柱子
		for r := i + 1; r < len(height); r++ {
			rightMax = maxT(rightMax, height[r])
		}

		minTrap := minT(leftMax, rightMax)
		if minTrap > height[i] {
			sum += minTrap - height[i]
		}
	}
	return sum
}

func maxT(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minT(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
