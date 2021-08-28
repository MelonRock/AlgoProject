package tree

func numTrees(n int) int {
	memo := make([]int, n+1)
	return helper(n, &memo)
}

func helper(n int, memo *[]int) int {
	if n == 1 || n == 0 {
		return 1
	}
	// 备忘录
	if (*memo)[n] > 0 {
		return (*memo)[n]
	}
	// 记录以i为根节点，i左子树的组合数 * n-i-1右子树的组合数
	count := 0
	// dp解决，
	/**
	左子树用掉 j 个，则右子树用掉 i-j-1 个，能构建出 dp[j] * dp[i-j-1] 种不同的BST
	*/
	for i := 0; i <= n-1; i++ {
		count += helper(i, memo) * helper(n-i-1, memo)
	}
	(*memo)[n] = count
	return count
}
