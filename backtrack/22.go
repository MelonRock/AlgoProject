package backtrack

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs22 func(left, right int, s string)
	dfs22 = func(left, right int, s string) {
		// 只有左括号多余右括号才有可能匹配
		if left < right || right > n || left > n {
			return
		}
		if left == n && right == n {
			res = append(res, s)
			return
		}
		dfs22(left+1, right, s+"(")
		dfs22(left, right+1, s+")")
	}

	dfs22(0, 0, "")
	return res
}
