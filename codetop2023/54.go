package codetop2023

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := make([]int, 0)
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top < bottom && left < right {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- {
			res = append(res, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			res = append(res, matrix[i][left])
		}
		top++
		bottom--
		left++
		right--
	}
	// 剩下一行
	if top == bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
	} else if left == right {
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
	}
	return res
}
