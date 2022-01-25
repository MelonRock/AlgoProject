package codetop

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 如果找到了 节点p或者q  就返回
	if root == p || root == q {
		return root
	}
	// 后序遍历，自底向上
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果left 和 right都不为空，说明此时root就是最近公共节点
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	// left为空，right不为空，就返回right，说明目标节点是通过right返回的
	if right != nil {
		return right
	}
	return nil
}
