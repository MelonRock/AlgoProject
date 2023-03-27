package codetop2023

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	target := n - k
	quickSort(nums, 0, n-1)
	return nums[target]
}

// 快排
func quickSort(nums []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	p := partition(nums, lo, hi)
	quickSort(nums, lo, p-1)
	quickSort(nums, p+1, hi)
}

func partition(nums []int, lo int, hi int) int {
	pivot := nums[lo]
	i, j := lo, hi
	for i < j {
		for i < j && nums[i] <= pivot {
			i++
		}
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 将基准数交换至两子数组的分界线
	nums[i], nums[lo] = nums[lo], nums[i]
	// 返回基准数的索引
	return i
}
