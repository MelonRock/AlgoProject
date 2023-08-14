package codetop2023

func maximalSquare(matrix [][]byte) int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return 0
	}
	height := len(matrix)
	width := len(matrix[0])
	maxSide := 0
	// dp(i + 1, j + 1) 是以 matrix(i, j) 为右下角的正方形的最大边长
	dp := make([][]int, height+1)
	for i := 0; i <= height; i++ {
		dp[i] = make([]int, width+1)
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			// 遇到1的正方形
			if matrix[row][col] == '1' {
				//  min(上, 左, 左上) + 1
				dp[row+1][col+1] = min221(min221(dp[row+1][col], dp[row][col+1]), dp[row][col]) + 1
				maxSide = max221(maxSide, dp[row+1][col+1])
			}
		}
	}
	return maxSide * maxSide

}

func min221(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max221(a, b int) int {
	if a > b {
		return a
	}
	return b
}
