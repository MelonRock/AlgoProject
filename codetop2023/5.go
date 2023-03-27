package codetop2023

func longestPalindrome(s string) string {
	res := ""
	// 中心扩散法
	for i := 0; i < len(s); i++ {
		// 以i为中心
		str := palindrome(s, i, i)
		if len(str) > len(res) {
			res = str
		}
		// 以i+1为中心
		str = palindrome(s, i, i+1)
		if len(str) > len(res) {
			res = str
		}
	}
	return res
}

func palindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
