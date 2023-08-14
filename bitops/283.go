package bitops

func moveZeroes(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		// 当指针i遇到不为0的，移动到j的位置
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < n; i++ {
		nums[i] = 0
	}
}
