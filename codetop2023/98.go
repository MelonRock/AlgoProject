package codetop2023

import "math"

func isValidBST(root *TreeNode) bool {
	var pre = math.MinInt64
	return helper98(root, &pre)
}

// 中序遍历

func helper98(root *TreeNode, pre *int) bool {
	if root == nil {
		return true
	}
	if !helper98(root.Left, pre) {
		return false
	}
	// 比较当前节点必须大于前一个节点
	if root.Val <= *pre {
		return false
	}
	*pre = root.Val
	return helper98(root.Right, pre)
}
