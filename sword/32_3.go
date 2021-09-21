package sword

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder32_3(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var isReverse bool
	for len(queue) > 0 {
		size := len(queue)
		list := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			if !isReverse {
				list[i] = node.Val
			} else {
				list[size-i-1] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		isReverse = !isReverse
		res = append(res, list)
	}
	return res
}
