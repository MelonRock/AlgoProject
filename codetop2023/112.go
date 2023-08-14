package codetop2023

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs112(root, targetSum, 0)
}

func dfs112(root *TreeNode, target int, sum int) bool {
	if root == nil {
		return false
	}
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		if sum == target {
			return true
		}
	}
	return dfs112(root.Left, target, sum) || dfs112(root.Right, target, sum)
}
