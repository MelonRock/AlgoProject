package sword

func replaceSpace(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if b[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}
