package codetop2023

// https://leetcode.cn/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	// 找到分界 A[i] < A[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 { // 不是最后一个排列
		// 找到[j,end)中尽可能小的数来交换
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 此时[j, end)是降序，逆置变升序
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}
