package hot100

import "math/rand"

func findKthLargest(nums []int, k int) int {
	nums = quickSort(nums)
	return nums[len(nums)-k]
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivotIdx := rand.Intn(len(nums))
	pivot := nums[pivotIdx]
	left, right := make([]int, 0), make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i == pivotIdx {
			continue
		}
		if nums[i] <= pivot {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	result := append(left, pivot)
	result = append(result, right...)
	return result
}
