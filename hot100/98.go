package hot100

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt32
	return dfs98(root, &pre)
}

func dfs98(root *TreeNode, pre *int) bool {
	if root == nil {
		return true
	}
	if !dfs98(root.Left, pre) {
		return false
	}
	if root.Val <= *pre {
		return false
	}
	*pre = root.Val
	return dfs98(root.Right, pre)
}
