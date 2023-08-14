package hot100

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	visit := make([]bool, len(nums))
	path := make([]int, 0)
	dfs46(nums, path, visit, &res)
	return res
}

func dfs46(nums []int, path []int, visit []bool, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visit[i] == true {
			continue
		}
		path = append(path, nums[i])
		visit[i] = true
		dfs46(nums, path, visit, res)
		visit[i] = false
		path = path[:len(path)-1]
	}
}
