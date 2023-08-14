package codetop2023

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var dfs124 func(root *TreeNode) int
	dfs124 = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs124(root.Left)
		right := dfs124(root.Right)
		// left -> root -> right作为路径与已经计算过的历史最大值做比较
		curMaxSum := left + root.Val + right
		maxSum = maxInt(maxSum, curMaxSum)
		// 返回经过root的单边最大分支给当前root的父节点使用
		retMaxSum := root.Val + maxInt(left, right)
		return maxInt(retMaxSum, 0)
	}
	dfs124(root)
	return maxSum
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
