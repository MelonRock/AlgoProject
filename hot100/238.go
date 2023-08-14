package hot100

func productExceptSelf(nums []int) []int {
	// 前缀积
	n := len(nums)
	ans := make([]int, n)

	// 左侧所有元素之积
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	// 再求右侧所有元素之积, 并与左侧所有元素之积相乘
	r := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] = ans[i] * r
		r = r * nums[i]
	}
	return ans
}
