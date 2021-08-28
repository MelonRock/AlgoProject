package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	traverse538(root, &sum)
	return root
}

func traverse538(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	// 后序遍历，从大到小
	traverse538(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	traverse538(root.Left, sum)
}
