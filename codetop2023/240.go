package codetop2023

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		l, r := 0, m-1
		for l < r {
			mid := (l + r + 1) / 2
			if matrix[i][mid] <= target {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if matrix[i][r] == target {
			return true
		}
	}
	return false
}
