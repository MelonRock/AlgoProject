package sword

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 先序遍历，然后交换左右节点
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 每个root节点需要交换它的左右孩子节点
	swap(root)
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

func swap(root *TreeNode) {
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
}
