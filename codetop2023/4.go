package codetop2023

func rightSideView(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == l-1 {
				res = append(res, node.Val)
			}
		}
	}
	return res
}
