package codetop2023

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	visit := make([]bool, len(nums))
	backtrack(nums, track, visit, &res)
	return res
}

func backtrack(nums []int, track []int, visit []bool, res *[][]int) {
	// base case,到达叶子节点，开始收集
	if len(nums) == len(track) {
		t := make([]int, len(track))
		copy(t, track)
		*res = append(*res, t)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visit[i] == true {
			continue
		}
		// 做选择
		track = append(track, nums[i])
		visit[i] = true
		// 回溯进入下一层
		backtrack(nums, track, visit, res)
		// 取消选择
		track = track[0 : len(track)-1]
		visit[i] = false
	}
}
