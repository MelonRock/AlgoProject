package hot100

// 使用数组来记录有0的行和列，第二次遍历时，对数组行列判断是否存在0，若存在，则将该位置置0。
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	row := make([]bool, m)
	col := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}
