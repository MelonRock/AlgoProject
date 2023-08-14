package hot100

// https://leetcode.cn/problems/path-sum-iii/solutions/109711/custerxue-xi-bi-ji-er-cha-shu-de-di-gui-he-dfs-by-/?envType=study-plan-v2&envId=top-100-liked

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return dfs437(root, targetSum) + dfs437(root.Left, targetSum) + dfs437(root.Right, targetSum)
}

func dfs437(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.Val == sum {
		count++
	}
	count += dfs437(root.Left, sum-root.Val)
	count += dfs437(root.Right, sum-root.Val)
	return count
}
