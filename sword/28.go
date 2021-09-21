package sword

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur28(root.Left, root.Right)
}

func recur28(L *TreeNode, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}
	if L == nil || R == nil || L.Val != R.Val {
		return false
	}
	return recur28(L.Left, R.Right) && recur28(L.Right, R.Left)
}
