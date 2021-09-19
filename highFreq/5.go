package highFreq

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以s[i]为中心的最长回文子串
		s1 := palindrome(s, i, i)
		if len(res) < len(s1) {
			res = s1
		}
		// 以s[i]和s[i+1]为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l, r int) string {
	// 防止索引越界
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 向两边展开
		l--
		r++
	}
	return s[l+1 : r]
}
