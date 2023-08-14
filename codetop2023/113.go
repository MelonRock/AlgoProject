package codetop2023

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	dfs113(root, targetSum, []int{}, &res)
	return res
}

func dfs113(root *TreeNode, sum int, arr []int, res *[][]int) {
	if root == nil {
		return
	}
	// 当前节点加入结果
	sum -= root.Val
	arr = append(arr, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		// 到达叶子节点
		path := make([]int, len(arr))
		copy(path, arr)
		*res = append(*res, path)
	}
	dfs113(root.Left, sum, arr, res)
	dfs113(root.Right, sum, arr, res)
	arr = arr[:len(arr)-1]
	return
}
