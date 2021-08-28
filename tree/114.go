package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 定义将以root为根的树拉平为链表
func flatten(root *TreeNode) {
	// base case
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	// 后序遍历，左右子树已经被拉平成一条链表
	left, right := root.Left, root.Right
	// 将左子树作为右子树
	root.Left = nil
	root.Right = left
	// 将原先的右子树接到当前右子树的末端
	p := root
	for p.Right != nil { // 寻找新的右子树的末端
		p = p.Right
	}
	p.Right = right
}
