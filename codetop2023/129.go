package codetop2023

func sumNumbers(root *TreeNode) int {
	return dfs129(root, 0)
}

func dfs129(root *TreeNode, cur int) int {
	if root == nil {
		return 0
	}
	// 当前cur的值需要更新
	cur = cur*10 + root.Val
	// 遍历到叶子节点
	if root.Left == nil && root.Right == nil {
		return cur
	}
	return dfs129(root.Left, cur) + dfs129(root.Right, cur)
}
