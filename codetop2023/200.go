package codetop2023

func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				// 遇到岛屿开始搜索
				dfsIslands(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfsIslands(grid [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfsIslands(grid, i-1, j)
	dfsIslands(grid, i+1, j)
	dfsIslands(grid, i, j-1)
	dfsIslands(grid, i, j+1)
}
