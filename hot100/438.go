package hot100

func findAnagrams(s string, p string) []int {
	win := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left := 0
	right := 0
	match := 0
	ans := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}
		// 当窗口长度大于字符串长度，锁紧窗口
		for right-left >= len(p) {
			// 当窗口长度等于字符串长度，且匹配完成
			if right-left == len(p) && match == len(need) {
				ans = append(ans, left)
			}
			d := s[left]
			left++
			if need[d] != 0 {
				if win[d] == need[d] {
					match--
				}
				win[d]--
			}
		}
	}
	return ans
}
