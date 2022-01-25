package sword

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	dict := map[byte]bool{}
	str := []byte(s)
	left := 0
	for right, val := range str {
		for dict[val] {
			_, ok := dict[str[left]]
			if ok {
				dict[str[left]] = false
			}
			left++
		}
		dict[val] = true
		maxLen = max48(maxLen, right-left+1)
	}
	return maxLen
}

func max48(a, b int) int {
	if a > b {
		return a
	}
	return b
}
