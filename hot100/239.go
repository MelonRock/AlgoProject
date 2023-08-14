package hot100

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	queue := []int{}
	res := make([]int, 0, n-k+1)

	for i, num := range nums {
		// 出界的下标删除
		if i >= k && queue[0] <= i-k {
			queue = queue[1:]
		}

		// 将下标i加入到队列，并保证队列是递减的
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= num {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// 将窗口最大元素加入结果集
		if i >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}

	return res
}
