package codetop2023

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs101(root.Left, root.Right)
}

func dfs101(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return dfs101(left.Left, right.Right) &&
		dfs101(left.Right, right.Left)
}
