package sword

import "math/rand"

var kk int

func getLeastNumbers(arr []int, k int) []int {
	kk = k
	quickSort(arr, 0, len(arr)-1)
	return arr[:k]
}

func quickSort(nums []int, left, right int) bool {
	if left > right {
		return false
	}
	key := rand.Int()%(right-left+1) + left
	nums[left], nums[key] = nums[key], nums[left]
	i, j, pivot := left, right, nums[left]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	if kk <= left {
		return true
	}
	if quickSort(nums, left, i-1) {
		return true
	}
	if quickSort(nums, i+1, right) {
		return true
	}
	return false
}
