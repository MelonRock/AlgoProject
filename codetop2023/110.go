package codetop2023

func isBalanced(root *TreeNode) bool {
	return recur(root) != -1
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	if right == -1 {
		return -1
	}

	if abs110(left-right) < 2 {
		// 以节点 root 为根节点的子树的最大高度
		return max110(left, right) + 1
	}
	return -1
}

func abs110(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func max110(a, b int) int {
	if a > b {
		return a
	}
	return b
}
