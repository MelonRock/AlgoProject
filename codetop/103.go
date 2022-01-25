package codetop

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	isReverse := false
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	for len(queue) != 0 {
		n := len(queue)
		level := make([]int, n)
		for i := 0; i < n; i++ {
			node := queue[i]
			if isReverse {
				level[n-i-1] = node.Val
			} else {
				level[i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		res = append(res, level)
		isReverse = !isReverse
	}
	return res
}
