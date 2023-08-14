package hot100

import "math"

func minWindow(s string, t string) string {
	win := make(map[byte]int)
	need := make(map[byte]int)
	match := 0
	left, right := 0, 0
	start, end := 0, 0
	min := math.MaxInt
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	var c byte
	for right < len(s) {
		c = s[right]
		right++
		if need[c] != 0 {
			win[c]++
			if need[c] == win[c] {
				match++
			}
		}
		// 开始收缩
		for match == len(need) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c = s[left]
			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
			left++
		}
	}
	if min == math.MaxInt {
		return ""
	}
	return s[start:end]
}
