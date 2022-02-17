package tree

// 非递归中序遍历二叉树
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := make([]*TreeNode, 0)
	if root == nil {
		return res
	}

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}
	return res
}
