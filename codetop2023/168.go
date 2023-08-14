package codetop2023

func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber > 0 {
		columnNumber--
		mod := columnNumber % 26
		res += string(rune('A' + mod))
		columnNumber = columnNumber / 26
	}
	i := 0
	j := len(res) - 1
	b := []byte(res)
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
