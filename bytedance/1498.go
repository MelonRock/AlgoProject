package bytedance

import "sort"

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	mod := 1000000007
	tmp := make([]int, n)
	tmp[0] = 1
	// 长度为n>1时的数组的子序列组合数为2^n
	for i := 1; i < n; i++ {
		// << 1 左移一位等于 *2
		// >> 1 右移一位等于 /2
		tmp[i] = (tmp[i-1] << 1) % mod
	}
	res := 0
	l, r := 0, n-1
	for l <= r {
		// 表示以l为左边界，l+1至r之间任一数字作为右边界，都满足上述条件，共计组合个数
		if nums[l]+nums[r] > target {
			r--
		} else {
			res = (res + tmp[r-l]) % mod
			l++
		}
	}
	return res
}
