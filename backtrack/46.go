package backtrack

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	visit := make([]bool, len(nums))
	helper(nums, track, visit, &res)
	return res
}

func helper(nums []int, track []int, visit []bool, res *[][]int) {
	if len(track) == len(nums) {
		t := make([]int, len(track))
		copy(t, track)
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visit[i] == true {
			continue
		}
		// 做选择， 加入路径
		track = append(track, nums[i])
		visit[i] = true
		// 进入下一个决策树
		helper(nums, track, visit, res)
		visit[i] = false
		track = track[0 : len(track)-1]
	}
}
