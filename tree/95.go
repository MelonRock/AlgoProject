package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// ***复习***
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	// 构造闭区间 [1, n] 组成的 BST
	return helper1(1, n)
}

func helper1(start, end int) []*TreeNode {
	if start > end {
		// base case
		return []*TreeNode{nil}
	}

	res := []*TreeNode{}
	// 1、穷举 root 节点的所有可能。
	for i := start; i <= end; i++ {
		// 2、递归构造出左右子树的所有合法 BST。
		left := helper1(start, i-1)
		right := helper1(i+1, end)
		// 3、给 root 节点穷举所有左右子树的组合。
		for _, l := range left {
			for _, r := range right {
				res = append(res, &TreeNode{i, l, r})
			}
		}
	}
	return res
}
