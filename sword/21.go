package sword

// 快慢指针, 如果nums[j]为奇数则与nums[i]进行交换并增加i的索引，循环直到j到达数组尾部，退出循环返回结果。
func exchange(nums []int) []int {
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j]%2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return nums
}
