package codetop2023

// 后序遍历，迭代版
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// 当前结果为根右左，需要反转
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

func postOrderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var postOrder func(root *TreeNode)
	postOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postOrder(root.Left)
		postOrder(root.Right)
		res = append(res, root.Val)
	}
	postOrder(root)
	return res
}
