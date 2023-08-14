package hot100

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	win := make(map[byte]int)
	res := 0
	j := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		win[c]++

		for win[c] > 1 {
			win[s[j]]--
			j++
		}
		res = max3(res, i-j+1)
	}
	return res
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
