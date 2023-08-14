package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	dfs543(root, &ans)
	return ans
}

func dfs543(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := dfs543(root.Left, ans)
	right := dfs543(root.Right, ans)
	// 经过 root，左右子树的最大深度之和
	*ans = max543(*ans, left+right)
	// 返回以该节点为根的二叉树的最大高度
	return max543(left, right) + 1
}

func max543(a, b int) int {
	if a > b {
		return a
	}
	return b
}
