package array

func removeDuplicateLetters(s string) string {
	time := [26]int{}
	var stack []byte
	inStack := [26]bool{}
	// 遍历数组，把字符出现次数存储起来
	for _, ch := range s {
		time[ch-'a']++
	}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			// 该位置的字母没有出现过
			for len(stack) > 0 && ch < stack[len(stack)-1] && time[stack[len(stack)-1]-'a'] > 0 {
				// 修改出现的状态
				inStack[stack[len(stack)-1]-'a'] = false
				// 栈顶元素弹出
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		time[ch-'a']--
	}
	return string(stack)
}
