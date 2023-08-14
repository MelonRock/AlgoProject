package codetop2023

func titleToNumber(columnTitle string) int {
	b := []byte(columnTitle)
	n := len(b)
	ans := 0
	for i := 0; i < n; i++ {
		c := b[i]
		ans = ans*26 + int(c-'A'+1)
	}
	return ans
}
