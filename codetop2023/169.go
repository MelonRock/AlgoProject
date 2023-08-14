package codetop2023

import (
	"math/rand"
	"sort"
	"time"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func quickSort169(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(nums) - 1)
	pivot := nums[index]
	left, right := make([]int, 0), make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i == index {
			continue
		}
		if nums[i] <= pivot {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	// 继续递归左右结果
	left = quickSort169(left)
	right = quickSort169(right)

	result := append(left, pivot)
	result = append(result, right...)
	return result
}

func mergeSort169(nums []int) []int {
	// 递归终止条件
	if len(nums) < 2 {
		return nums
	}
	// 每次分割成左右两部分，直到不能分割
	mid := len(nums) / 2
	left := mergeSort169(nums[:mid])
	right := mergeSort169(nums[mid:])
	return merge169(left, right)
}

func merge169(left, right []int) []int {
	i, j := 0, 0
	var res []int
	// 遍历两个子数组，排序加入新的结果
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			res = append(res, right[j])
			j++
		} else {
			res = append(res, left[i])
			i++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[i:]...)
	return res
}
