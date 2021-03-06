package array

// 3.无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	win := make(map[byte]int)
	left := 0
	right := 0
	ans := 1
	for right < len(s) {
		c := s[right]
		right++
		win[c]++
		// 窗口缩小,直到没有重复元素
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}
		ans = max(right-left, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
