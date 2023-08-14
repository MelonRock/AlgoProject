package codetop2023

// https://leetcode.cn/problems/diameter-of-binary-tree/solution/dfsdi-gui-jie-jue-jian-dan-yi-dong-xiao-dta1o/
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	// 该二叉树左子树的最大高度+右子树的最大高度
	dfs543(root, &ans)
	return ans
}

// 找出以 root 为根节点的二叉树的最大深度
func dfs543(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	l := dfs543(root.Left, max)
	r := dfs543(root.Right, max)
	// 经过 root，左右子树的最大深度之和
	*max = max543(*max, l+r)
	// 返回以该节点为根的二叉树的最大高度
	return max543(l, r) + 1
}

func max543(a, b int) int {
	if a > b {
		return a
	}
	return b
}
