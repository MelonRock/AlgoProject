package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs94 func(root *TreeNode)
	dfs94 = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs94(root.Left)
		res = append(res, root.Val)
		dfs94(root.Right)
	}
	dfs94(root)
	return res
}
