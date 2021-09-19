package bitops

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	target := n - k
	array := quickSort(nums, 0, n)
	return array[target]
}

func partition(nums []int, left, right int) int {
	povit := left
	index := povit + 1
	// i 是当前遍历元素的索引， index是记录分界点
	for i := index; i <= len(nums); i++ {
		if nums[i] < nums[povit] {
			nums[i], nums[povit] = nums[povit], nums[i]
			index++
		}
	}
	nums[povit], nums[index-1] = nums[index-1], nums[povit]
	return index - 1
}

func quickSort(nums []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(nums, 0, len(nums))
		quickSort(nums, 0, partitionIndex-1)
		quickSort(nums, partitionIndex+1, right)
	}
	return nums
}
