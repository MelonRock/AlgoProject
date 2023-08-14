package codetop2023

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var dfs39 func(start int, path []int, sum int)

	dfs39 = func(start int, path []int, sum int) {
		if sum >= target {
			// 找到组合
			if sum == target {
				temp := make([]int, len(path))
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}
		// start
		for i := start; i < len(candidates); i++ {
			// 做选择
			path = append(path, candidates[i])
			// 基于i做选择，下次就不会选到i左边的数,达到去重
			dfs39(i, path, sum+candidates[i])
			// 取消选择
			path = path[:len(path)-1]
		}
		return
	}
	dfs39(0, []int{}, 0)
	return res
}
