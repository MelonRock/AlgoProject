package sword

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder32_2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		n := len(stack)
		var v []int
		for i := 0; i < n; i++ {
			node := stack[0]
			v = append(v, node.Val)
			stack = stack[1:]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		res = append(res, v)
	}
	return res
}
