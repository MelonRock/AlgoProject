package codetop2023

// 中序遍历，迭代法
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root == nil {
		return res
	}
	for len(stack) > 0 || root != nil {
		if root != nil {
			// 不断往左子树方向走
			stack = append(stack, root)
			root = root.Left
		} else {
			// 左子树走到底，弹出元素，然后开始走右子树
			node := stack[len(stack)-1]
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}

// 中序遍历，递归版
func inOrderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return res
}

// 144 145
