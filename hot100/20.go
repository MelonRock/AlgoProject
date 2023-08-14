package hot100

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		']': '[',
		')': '(',
		'{': '}',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if pairs[b] > 0 {
			if len(stack) > 0 && pairs[b] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, b)
		}
	}
	return len(stack) == 1
}
