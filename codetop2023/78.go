package codetop2023

func subsets(nums []int) [][]int {
	track := make([]int, 0)
	res := make([][]int, 0)
	dfs(nums, 0, &res, track)
	return res
}

func dfs(nums []int, start int, res *[][]int, track []int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)
	// 使用 start 参数控制树枝的生长避免产生重复的子集
	for i := start; i < len(nums); i++ {
		// 做选择
		track = append(track, nums[i])
		// 进入下一层
		dfs(nums, i+1, res, track)
		// 回退
		track = track[:len(track)-1]
	}
	return
}
