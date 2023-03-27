package backtrack

func findSubsequences(nums []int) (ans [][]int) {
	var helper func(path []int, start int)
	helper = func(path []int, start int) {
		if len(path) > 1 {
			dest := make([]int, len(path))
			copy(dest, path) // 注意：这里要拷贝一份，防止同一份path被多次修改
			ans = append(ans, dest)
		}
		// 收集结果
		set := map[int]bool{} // 本层共用的map，用于保证本层添加的元素被去重
		// 确保遍历到每个子集
		for i := start; i < len(nums); i++ {
			// 剪枝：本层去重
			if set[nums[i]] {
				continue
			}
			// 剪枝：不符合递增要求
			if len(path) > 0 && path[len(path)-1] > nums[i] {
				continue
			}
			// 维护set
			set[nums[i]] = true
			// 递归下一层
			path = append(path, nums[i])
			helper(path, i+1)
		}
	}
	helper([]int{}, 0)
	return
}
