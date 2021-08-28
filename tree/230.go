package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	res := 0
	rank := 0
	traverse2(root, k, &res, &rank)
	return res
}

func traverse2(root *TreeNode, k int, res *int, rank *int) {
	if root == nil {
		return
	}
	traverse2(root.Left, k, res, rank)
	// 当前root所在位置
	*rank++
	// 找到第k小的元素
	if k == *rank {
		*res = root.Val
		return
	}
	traverse2(root.Right, k, res, rank)
}
