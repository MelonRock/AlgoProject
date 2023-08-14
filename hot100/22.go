package hot100

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	s := ""
	var dfs22 func(left, right int, s string)
	dfs22 = func(left, right int, s string) {
		if left < right || left > n || right > n {
			return
		}
		if left == n && right == n {
			res = append(res, s)
			return
		}
		dfs22(left+1, right, s+"(")
		dfs22(left, right+1, s+")")
	}
	dfs22(0, 0, s)
	return res
}
