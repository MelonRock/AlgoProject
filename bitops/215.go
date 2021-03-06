package bitops

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	target := n - k
	quickSort(nums, 0, n-1)
	return nums[target]
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	index := left
	// i 是当前遍历元素的索引， index是记录分界点
	for i := left + 1; i <= right; i++ {
		if nums[i] < pivot {
			index++
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	nums[left], nums[index] = nums[index], nums[left]
	return index
}

func quickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	partitionIndex := partition(nums, left, right)
	quickSort(nums, left, partitionIndex-1)
	quickSort(nums, partitionIndex+1, right)
}
