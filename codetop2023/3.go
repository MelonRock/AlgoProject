package codetop2023

// 无重复字符的最长子串 滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	windows := make(map[byte]int)
	j := 0
	res := 0
	for i := 0; i < len(s); {
		c := s[i]
		windows[c] += 1
		i++
		for windows[c] > 1 {
			windows[s[j]] -= 1
			j++
		}
		res = max(res, i-j)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
