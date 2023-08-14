package hot100

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var dfs39 func(start int, path []int, sum int)
	dfs39 = func(start int, path []int, sum int) {
		if sum >= target {
			if sum == target {
				temp := make([]int, len(path))
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs39(i, path, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	dfs39(0, []int{}, 0)
	return res
}
