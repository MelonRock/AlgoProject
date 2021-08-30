package array

import (
	"math"
)

func minWindow(s string, t string) string {
	// 保存滑动窗口字符集
	win := make(map[byte]int)
	// 保存需要的字符集
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	// 窗口
	left := 0
	right := 0
	// match匹配次数
	match := 0
	start := 0           // 记录最小覆盖子串的起始索引
	end := 0             // 记录最小覆盖子串的结束索引
	min := math.MaxInt64 // 所选的字符长度
	var c byte
	for right < len(s) {
		c = s[right]
		right++
		// 在需要的字符集里，添加到窗口字符集
		if need[c] != 0 {
			win[c]++
			// 如果当前字符数的数量与需要的数量相同
			if win[c] == need[c] {
				match++
			}
		}

		// 当所有字符数量匹配上，开始收缩
		for match == len(need) {
			// 记录符合条件的窗口索引
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c = s[left]
			left++
			// 左指针的字符不在需要的字符集
			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}
