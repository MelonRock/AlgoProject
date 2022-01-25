package codetop

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	visit := make([]bool, len(nums))
	backtrack(nums, track, visit, &res)
	return res
}

func backtrack(nums []int, track []int, visit []bool, res *[][]int) {
	if len(nums) == len(track) {
		t := make([]int, len(nums))
		copy(t, track)
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visit[i] == true {
			continue
		}
		track = append(track, nums[i])
		visit[i] = true
		backtrack(nums, track, visit, res)
		visit[i] = false
		track = track[0 : len(track)-1]
	}
}
