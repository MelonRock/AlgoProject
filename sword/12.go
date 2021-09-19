package sword

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果在数组中找到第一个数，就执行下一个，否则返回false
			if search(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, i, j, k int, word string) bool {
	// 找到最后一个数，返回true
	if k == len(word) {
		return true
	}
	// i，j的约束条件
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}
	// 进入dfs， 先把正在遍历的该值重新赋值，如果在该值周围都搜索不到目标字符
	// 则再把该值还原，如果在数组找到第一个字符，则进入下一个字符超找
	if board[i][j] == word[k] {
		temp := board[i][j]
		board[i][j] = ' '
		// 如果if进入成功，说明找到该字符，然后进行下一个字符搜索，直到所有搜索
		// 都成功，即k == len(word) -1的大小时，会返回true，进入该条件语句，
		// 然后返回函数true值
		if search(board, i, j+1, k+1, word) ||
			search(board, i, j-1, k+1, word) ||
			search(board, i+1, j, k+1, word) ||
			search(board, i-1, j, k+1, word) {
			return true
		} else {
			board[i][j] = temp
		}
	}
	return false
}
