package codetop

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	target := n - k
	quickSort(nums, 0, n-1)
	return nums[target]
}

func partition(nums []int, left int, right int) int {
	pivot := nums[left]
	index := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < pivot {
			index++
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	nums[left], nums[index] = nums[index], nums[left]
	return index
}

func quickSort(nums []int, left int, right int) {
	if left > right {
		return
	}
	partitions := partition(nums, left, right)
	quickSort(nums, left, partitions-1)
	quickSort(nums, partitions+1, right)
}
