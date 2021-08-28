package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBst(root, nil, nil)
}

// 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val
func isValidBst(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}
	// 若 root.val 不符合 max 和 min 的限制，说明不是合法 BST
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	// 限定左子树的最大值是 root.val，右子树的最小值是 root.val
	return isValidBst(root.Left, min, root) && isValidBst(root.Right, root, max)
}
