package codetop2023

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
		v := make([]int, n)
		for i := 0; i < n; i++ {
			node := queue[i]
			if isReverse {
				v[n-i-1] = node.Val
			} else {
				v[i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		isReverse = !isReverse
		queue = queue[n:]
		res = append(res, v)
	}
	return res
}
