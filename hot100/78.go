package hot100

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	dfs78(nums, 0, []int{}, &res)
	return res
}

func dfs78(nums []int, start int, path []int, res *[][]int) {
	{
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs78(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
