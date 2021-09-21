package sword

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 先序遍历树 A 中的每个节点 nA
	return (A != nil && B != nil) && (isSubStructure(A.Left, B) ||
		isSubStructure(A.Right, B) || recur(A, B))
}

// 判断树 A 中 以 nA 为根节点的子树 是否包含树 B
func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
