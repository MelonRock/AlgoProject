package sword

func verifyPostorder(postorder []int) bool {
	// [ 左子树 | 右子树 | 根节点 ]
	return recur33(postorder, 0, len(postorder)-1)
}

func recur33(postorder []int, i int, j int) bool {
	if i >= j {
		return true
	}
	// 寻找第一个大于根节点的元素，
	// 左子树区间 [i,m−1]
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	// 右子树区间[m,j−1]、根节点索引 j
	for postorder[p] > postorder[j] {
		p++
	}
	// p == j判断是否二叉搜索树, 判断左子树，判断右子树
	return p == j && recur33(postorder, i, m-1) && recur33(postorder, m, j-1)
}
