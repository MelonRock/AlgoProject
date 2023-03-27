package sort

import "math/rand"

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivotIdx := rand.Intn(len(nums) - 1)
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
	left = quickSort(left)
	right = quickSort(right)

	result := append(left, pivot)
	result = append(result, right...)
	return result
}
